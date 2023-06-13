package main

import (
	"context"
	"fmt"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {

	if req.Message == nil {
		return nil, fmt.Errorf("Message should not be nil")
	}
	
	id := GetUniqueID(req.Message.GetChat())

	message := &Input{
		Message:   req.Message.GetText(),
		Sender:    req.Message.GetSender(),
		Timestamp: req.Message.GetSendTime(),
	}

	if err := database.WriteToDatabase(ctx, id, message); err != nil {
		return nil, err
	}

	response := rpc.NewSendResponse()
	response.Code, response.Msg = 0, "success"
	return response, nil
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	
	chatID := GetUniqueID(req.GetChat())
	start := req.GetCursor()
	limit := req.GetLimit()
	end := start + int64(limit)


	messages, err := database.ReadFromDatabase(ctx, chatID, start, end, req.GetReverse())
	if err != nil {
		return nil, err
	}

	returnedMessages := make([]*rpc.Message, 0)
	var count int32 = 0
	var nextCursor int64 = 0
	hasMore := false
	for _, message := range messages {
		if count+1 > limit {
			hasMore = true
			nextCursor = end
			break
		}
		returnedMessages = append(returnedMessages, &rpc.Message{
			Chat:     req.GetChat(),
			Text:     message.Message,
			Sender:   message.Sender,
			SendTime: message.Timestamp,
		})
		count++
	}

	response := rpc.NewPullResponse()
	response.Messages, response.HasMore, response.NextCursor = returnedMessages, &hasMore, &nextCursor
	response.Code, response.Msg = 0, "success"
	return response, nil
}