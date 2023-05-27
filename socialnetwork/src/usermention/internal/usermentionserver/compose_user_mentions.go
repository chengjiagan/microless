package usermentionserver

import (
	"context"
	"microless/socialnetwork/proto"
	"microless/socialnetwork/proto/user"
	pb "microless/socialnetwork/proto/usermention"

	"golang.org/x/sync/errgroup"
)

func (s *UserMentionService) ComposeUserMentions(ctx context.Context, req *pb.ComposeUserMentionsRequest) (*pb.ComposeUserMentionsRespond, error) {
	mentions := make([]*proto.UserMention, len(req.Usernames))
	g, ctx := errgroup.WithContext(ctx)
	for i, name := range req.Usernames {
		i, name := i, name // see https://go.dev/doc/faq#closures_and_goroutines
		g.Go(func() error {
			userReq := &user.GetUserIdRequest{Username: name}
			userResp, err := s.userClient.GetUserId(ctx, userReq)
			if err != nil {
				s.logger.Warnw("Failed to get user_id from User service", "username", name, "err", err)
				return err
			}
			mentions[i] = &proto.UserMention{
				UserId:   userResp.UserId,
				Username: name,
			}
			return nil
		})
	}

	err := g.Wait()
	if err != nil {
		return nil, err
	}
	return &pb.ComposeUserMentionsRespond{UserMentions: mentions}, nil
}
