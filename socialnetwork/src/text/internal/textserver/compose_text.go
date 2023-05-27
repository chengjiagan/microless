package textserver

import (
	"context"
	pb "microless/socialnetwork/proto/text"
	"microless/socialnetwork/proto/urlshorten"
	"microless/socialnetwork/proto/usermention"
	"regexp"

	"golang.org/x/sync/errgroup"
)

func (s *TextService) ComposeText(ctx context.Context, req *pb.ComposeTextRequest) (*pb.ComposeTextRespond, error) {
	g, ctx := errgroup.WithContext(ctx)
	resp := &pb.ComposeTextRespond{Text: req.Text}

	// find user mentions in text and compose
	g.Go(func() error {
		re := regexp.MustCompile("@[a-zA-Z0-9-_]+")
		mentions := re.FindAllString(req.Text, -1)
		// contain no user mentions
		if len(mentions) == 0 {
			return nil
		}

		usernames := make([]string, len(mentions))
		for i, m := range mentions {
			usernames[i] = m[1:]
		}
		mentionReq := &usermention.ComposeUserMentionsRequest{Usernames: usernames}
		mentionResp, err := s.usermentionClient.ComposeUserMentions(ctx, mentionReq)
		if err != nil {
			s.logger.Warnw("Failed to compose user mentions from UserMention service", "err", err)
			return err
		}
		resp.UserMention = mentionResp.UserMentions
		return nil
	})

	// find urls in text and compose
	g.Go(func() error {
		re := regexp.MustCompile("(http://|https://)[a-zA-Z0-9_!~*'().&=+$%-]+")
		urls := re.FindAllString(req.Text, -1)
		// contain no urls
		if len(urls) == 0 {
			return nil
		}

		urlReq := &urlshorten.ComposeUrlsRequest{Urls: urls}
		urlResp, err := s.urlshortenClient.ComposeUrls(ctx, urlReq)
		if err != nil {
			s.logger.Warnw("Failed to shorten urls from UrlShorten service", "err", err)
			return err
		}
		resp.Urls = urlResp.Urls
		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}
	return resp, nil
}
