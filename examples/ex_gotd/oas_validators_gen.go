// Code generated by ogen, DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/json"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// No-op definition for keeping imports.
var (
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = sort.Ints
	_ = chi.Context{}
	_ = http.MethodGet
	_ = io.Copy
	_ = json.Marshal
	_ = bytes.NewReader
	_ = strconv.ParseInt
	_ = time.Time{}
	_ = conv.ToInt32
	_ = uuid.UUID{}
	_ = uri.PathEncoder{}
	_ = url.URL{}
	_ = math.Mod
	_ = validate.Int{}
	_ = ht.NewRequest
	_ = net.IP{}
)

func (s *Animation) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Duration)); err != nil {
			return err
		}
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Height)); err != nil {
			return err
		}
		_ = s.Height
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "height",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Thumb
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "thumb",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Width)); err != nil {
			return err
		}
		_ = s.Width
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "width",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AnswerCallbackQuery) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Text
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "text",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AnswerInlineQuery) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Results == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Results {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Results
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "results",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.SwitchPmParameter
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "switch_pm_parameter",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *AnswerShippingQuery) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.ShippingOptions {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.ShippingOptions
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "shipping_options",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Audio) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Duration)); err != nil {
			return err
		}
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Thumb
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "thumb",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *BotCommand) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    32,
			MaxLengthSet: true,
		}).Validate(string(s.Command)); err != nil {
			return err
		}
		_ = s.Command
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "command",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    3,
			MinLengthSet: true,
			MaxLength:    256,
			MaxLengthSet: true,
		}).Validate(string(s.Description)); err != nil {
			return err
		}
		_ = s.Description
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "description",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Chat) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Location
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "location",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PinnedMessage
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "pinned_message",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ChatLocation) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    64,
			MaxLengthSet: true,
		}).Validate(string(s.Address)); err != nil {
			return err
		}
		_ = s.Address
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "address",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CopyMessage) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *CreateNewStickerSet) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    64,
			MaxLengthSet: true,
		}).Validate(string(s.Name)); err != nil {
			return err
		}
		_ = s.Name
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "name",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    64,
			MaxLengthSet: true,
		}).Validate(string(s.Title)); err != nil {
			return err
		}
		_ = s.Title
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Document) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Thumb
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "thumb",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EditMessageCaption) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *EditMessageText) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    4096,
			MaxLengthSet: true,
		}).Validate(string(s.Text)); err != nil {
			return err
		}
		_ = s.Text
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "text",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Game) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Animation
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "animation",
			Error: err,
		})
	}
	if err := func() error {
		if s.Photo == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Photo {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Photo
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "photo",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Text
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "text",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUpdates) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Limit
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "limit",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *GetUserProfilePhotos) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Limit
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "limit",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Message) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Animation
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "animation",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Audio
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "audio",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if err := func() error {
		if s.Chat == nil {
			return fmt.Errorf("required, can't be nil")
		}
		_ = s.Chat
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "chat",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Document
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "document",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.ForwardFromChat
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "forward_from_chat",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Game
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "game",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.NewChatPhoto {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.NewChatPhoto
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "new_chat_photo",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PassportData
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "passport_data",
			Error: err,
		})
	}
	if err := func() error {
		var failures []validate.FieldError
		for i, elem := range s.Photo {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Photo
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "photo",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PinnedMessage
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "pinned_message",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Poll
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "poll",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.ReplyToMessage
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "reply_to_message",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.SenderChat
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sender_chat",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Sticker
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sticker",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Text
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "text",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Video
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.VideoNote
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "video_note",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Voice
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "voice",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.VoiceChatEnded
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "voice_chat_ended",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PassportData) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Data == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Data {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Data
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "data",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PhotoSize) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Height)); err != nil {
			return err
		}
		_ = s.Height
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "height",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Width)); err != nil {
			return err
		}
		_ = s.Width
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "width",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Poll) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Explanation
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "explanation",
			Error: err,
		})
	}
	if err := func() error {
		if s.Options == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Options {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Options
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "options",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    300,
			MaxLengthSet: true,
		}).Validate(string(s.Question)); err != nil {
			return err
		}
		_ = s.Question
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "question",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *PollOption) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    100,
			MaxLengthSet: true,
		}).Validate(string(s.Text)); err != nil {
			return err
		}
		_ = s.Text
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "text",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ResultMsg) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Result
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "result",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendAnimation) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Height
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "height",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Width
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "width",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendAudio) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendDocument) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendInvoice) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    255,
			MaxLengthSet: true,
		}).Validate(string(s.Description)); err != nil {
			return err
		}
		_ = s.Description
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "description",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PhotoHeight
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "photo_height",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.PhotoWidth
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "photo_width",
			Error: err,
		})
	}
	if err := func() error {
		if s.Prices == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Prices {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Prices
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "prices",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    32,
			MaxLengthSet: true,
		}).Validate(string(s.Title)); err != nil {
			return err
		}
		_ = s.Title
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendMediaGroup) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Media == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Media {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Media
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "media",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendMessage) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    4096,
			MaxLengthSet: true,
		}).Validate(string(s.Text)); err != nil {
			return err
		}
		_ = s.Text
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "text",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendPhoto) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendPoll) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Explanation
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "explanation",
			Error: err,
		})
	}
	if err := func() error {
		if s.Options == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Options {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Options
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "options",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    300,
			MaxLengthSet: true,
		}).Validate(string(s.Question)); err != nil {
			return err
		}
		_ = s.Question
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "question",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendVideo) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Height
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "height",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Width
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "width",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendVideoNote) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SendVoice) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Caption
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "caption",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SetChatAdministratorCustomTitle) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    0,
			MinLengthSet: false,
			MaxLength:    16,
			MaxLengthSet: true,
		}).Validate(string(s.CustomTitle)); err != nil {
			return err
		}
		_ = s.CustomTitle
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "custom_title",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SetChatDescription) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		_ = s.Description
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "description",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SetChatTitle) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.String{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    255,
			MaxLengthSet: true,
		}).Validate(string(s.Title)); err != nil {
			return err
		}
		_ = s.Title
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "title",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SetMyCommands) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Commands == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Commands {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Commands
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "commands",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *SetPassportDataErrors) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Errors == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Errors {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Errors
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "errors",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *ShippingOption) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if s.Prices == nil {
			return fmt.Errorf("required, can't be nil")
		}
		var failures []validate.FieldError
		for i, elem := range s.Prices {
			if err := func() error {
				_ = elem
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
			if len(failures) > 0 {
				return &validate.Error{Fields: failures}
			}
		}
		_ = s.Prices
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "prices",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Sticker) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Height)); err != nil {
			return err
		}
		_ = s.Height
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "height",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Thumb
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "thumb",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Width)); err != nil {
			return err
		}
		_ = s.Width
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "width",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Video) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Duration)); err != nil {
			return err
		}
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Height)); err != nil {
			return err
		}
		_ = s.Height
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "height",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Thumb
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "thumb",
			Error: err,
		})
	}
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Width)); err != nil {
			return err
		}
		_ = s.Width
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "width",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *VideoNote) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Duration)); err != nil {
			return err
		}
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if err := func() error {
		_ = s.Thumb
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "thumb",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *Voice) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Duration)); err != nil {
			return err
		}
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
func (s *VoiceChatEnded) Validate() error {
	var failures []validate.FieldError
	if err := func() error {
		if err := (validate.Int{
			MinSet:       true,
			Min:          0,
			MaxSet:       false,
			Max:          0,
			MinExclusive: true,
			MaxExclusive: false,
		}).Validate(int64(s.Duration)); err != nil {
			return err
		}
		_ = s.Duration
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "duration",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
