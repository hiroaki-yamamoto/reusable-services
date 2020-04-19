package sendfuncs_test

import (
	"context"
	"errors"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mailgun/mailgun-go/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	. "github.com/hiroaki-yamamoto/reusable-services/email/sendfuncs"
	"github.com/hiroaki-yamamoto/reusable-services/email/sendfuncs/mocks"
)

var _ = Describe("Mailgun", func() {
	var mailgunMock *mocks.MockMailgun
	var send *Mailgun
	var origin mailgun.Mailgun
	var ctx context.Context
	var cancel context.CancelFunc
	BeforeEach(func() {
		mailgunMock = mocks.NewMockMailgun(rootCtrl)
		send = NewMailgun(zap.NewNop(), "test", "test")
		origin = send.Mailgun
		send.Mailgun = mailgunMock
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	})
	AfterEach(func() {
		cancel()
	})
	Context("When send function gets succeeded", func() {
		var msg *mailgun.Message
		BeforeEach(func() {
			mailgunMock.EXPECT().NewMessage(
				gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
			).DoAndReturn(
				func(from, title, txt string, to ...string) *mailgun.Message {
					msg = origin.NewMessage(from, title, txt, to...)
					return msg
				},
			).Times(1)
			mailgunMock.EXPECT().Send(gomock.Any(), gomock.Any()).Return(
				"test response", "test_id", nil,
			).Times(1)
		})
		It("Shouldn't raise any errors", func() {
			err := send.Send(
				ctx, "testFrom@example.com", "testTo@example.com",
				"testTitle", "testBody", "<p>testBody</p>",
			)
			Expect(err).To(Succeed())
		})
	})
	Context("When send function gets failure", func() {
		var msg *mailgun.Message
		BeforeEach(func() {
			mailgunMock.EXPECT().NewMessage(
				gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
			).DoAndReturn(
				func(from, title, txt string, to ...string) *mailgun.Message {
					msg = origin.NewMessage(from, title, txt, to...)
					return msg
				},
			).Times(1)
			mailgunMock.EXPECT().Send(gomock.Any(), gomock.Any()).Return(
				"", "", errors.New("Test Error"),
			).Times(1)
		})
		It("Should raise an error.", func() {
			err := send.Send(
				ctx, "testFrom@example.com", "testTo@example.com",
				"testTitle", "testBody", "<p>testBody</p>",
			)
			Expect(err).To(MatchError(errors.New("Test Error")))
		})
	})
})
