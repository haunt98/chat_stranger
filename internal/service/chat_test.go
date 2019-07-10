package service

import (
	"testing"

	"github.com/1612180/chat_stranger/internal/mock/mock_repository"
	"github.com/1612180/chat_stranger/internal/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestChatService_Find(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("empty false", func(t *testing.T) {
		roomMock := mock_repository.NewMockRoomRepository(ctrl)
		memberMock := mock_repository.NewMockMemberRepo(ctrl)
		messageMock := mock_repository.NewMockMessageRepo(ctrl)

		roomMock.EXPECT().
			FindEmpty().
			Return(nil, false)
		roomMock.EXPECT().
			Create().
			Return(nil, false)

		chatService := ChatService{
			roomRepo:    roomMock,
			memberRepo:  memberMock,
			messageRepo: messageMock,
		}

		_, ok := chatService.Find(1, "empty")
		assert.Equal(t, false, ok)
	})

	t.Run("empty true", func(t *testing.T) {
		roomMock := mock_repository.NewMockRoomRepository(ctrl)
		memberMock := mock_repository.NewMockMemberRepo(ctrl)
		messageMock := mock_repository.NewMockMessageRepo(ctrl)

		roomMock.EXPECT().
			FindEmpty().
			Return(&model.Room{}, true)

		chatService := ChatService{
			roomRepo:    roomMock,
			memberRepo:  memberMock,
			messageRepo: messageMock,
		}

		_, ok := chatService.Find(1, "empty")
		assert.Equal(t, true, ok)
	})

	t.Run("empty true", func(t *testing.T) {
		roomMock := mock_repository.NewMockRoomRepository(ctrl)
		memberMock := mock_repository.NewMockMemberRepo(ctrl)
		messageMock := mock_repository.NewMockMessageRepo(ctrl)

		roomMock.EXPECT().
			FindEmpty().
			Return(nil, false)
		roomMock.EXPECT().
			Create().
			Return(&model.Room{}, true)

		chatService := ChatService{
			roomRepo:    roomMock,
			memberRepo:  memberMock,
			messageRepo: messageMock,
		}
		_, ok := chatService.Find(1, "empty")
		assert.Equal(t, true, ok)
	})
}
