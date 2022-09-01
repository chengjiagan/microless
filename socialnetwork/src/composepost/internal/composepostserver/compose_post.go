package composepostserver

import (
	"context"
	"microless/socialnetwork/proto"
	pb "microless/socialnetwork/proto/composepost"
	"microless/socialnetwork/proto/hometimeline"
	"microless/socialnetwork/proto/media"
	"microless/socialnetwork/proto/poststorage"
	"microless/socialnetwork/proto/text"
	"microless/socialnetwork/proto/user"
	"microless/socialnetwork/proto/usertimeline"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *ComposePostService) ComposePost(ctx context.Context, req *pb.ComposePostRequest) (*emptypb.Empty, error) {
	post := &proto.Post{
		PostType: req.PostType,
	}

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error { return s.composeText(gCtx, req.Text, post) })
	g.Go(func() error { return s.composeCreator(gCtx, req.UserId, req.Username, post) })
	g.Go(func() error { return s.composeMedia(gCtx, req.MediaTypes, req.MediaIds, post) })
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// upload post
	postReq := &poststorage.StorePostRequest{Post: post}
	postResp, err := s.poststorageClient.StorePost(ctx, postReq)
	if err != nil {
		s.logger.Errorw("Failed to upload post to PostStorage service", "err", err)
		return nil, err
	}

	// get user id mentioned in post
	mentions := make([]string, len(post.UserMentions))
	for i, user := range post.UserMentions {
		mentions[i] = user.UserId
	}

	// update timeline
	g, gCtx = errgroup.WithContext(ctx)
	g.Go(func() error { return s.uploadUserTimeline(gCtx, postResp.PostId, req.UserId, postResp.Timestamp) })
	g.Go(func() error {
		return s.uploadHomeTimeline(gCtx, postResp.PostId, req.UserId, postResp.Timestamp, mentions)
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// some helper functions

func (s *ComposePostService) composeText(ctx context.Context, postText string, post *proto.Post) error {
	req := &text.ComposeTextRequest{Text: postText}
	resp, err := s.textClient.ComposeText(ctx, req)
	if err != nil {
		s.logger.Errorw("Failed to compose text from Text service", "err", err)
		return err
	}

	post.Text = resp.Text
	post.Urls = resp.Urls
	post.UserMentions = resp.UserMention
	return nil
}

func (s *ComposePostService) composeCreator(ctx context.Context, userId string, username string, post *proto.Post) error {
	req := &user.ComposeCreatorWithUserIdRequest{UserId: userId, Username: username}
	resp, err := s.userClient.ComposeCreatorWithUserId(ctx, req)
	if err != nil {
		s.logger.Errorw("Failed to compose creator from User service", "err", err)
		return err
	}

	post.Creator = resp.Creator
	return nil
}

func (s *ComposePostService) composeMedia(ctx context.Context, mediaTypes []string, mediaIds []int64, post *proto.Post) error {
	req := &media.ComposeMediaRequest{MediaTypes: mediaTypes, MediaIds: mediaIds}
	resp, err := s.mediaClient.ComposeMedia(ctx, req)
	if err != nil {
		s.logger.Errorw("Failed to compose media from Media service", "err", err)
		return err
	}

	post.Media = resp.Media
	return nil
}

func (s *ComposePostService) uploadUserTimeline(ctx context.Context, postId, userId string, timestamp *timestamppb.Timestamp) error {
	req := &usertimeline.WriteUserTimelineRequest{PostId: postId, UserId: userId, Timestamp: timestamp}
	_, err := s.usertimelineClient.WriteUserTimeline(ctx, req)
	if err != nil {
		s.logger.Errorw("Failed to upload post to UserTimeline service", "err", err)
		return err
	}
	return nil
}

func (s *ComposePostService) uploadHomeTimeline(ctx context.Context, postId, userId string, timestamp *timestamppb.Timestamp, userMentionIds []string) error {
	req := &hometimeline.WriteHomeTimelineRequest{
		UserId:         userId,
		PostId:         postId,
		Timestamp:      timestamp,
		UserMentionsId: userMentionIds,
	}
	_, err := s.hometimelineClient.WriteHomeTimeline(ctx, req)
	if err != nil {
		s.logger.Errorw("Failed to upload post to HomeTimeline service", "err", err)
		return err
	}
	return nil
}
