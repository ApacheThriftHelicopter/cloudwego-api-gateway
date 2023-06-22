// Code generated by Validator v0.1.4. DO NOT EDIT.

package book

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *BaseResp) IsValid() error {
	return nil
}
func (p *Book) IsValid() error {
	return nil
}
func (p *CreateBookRequest) IsValid() error {
	if len(p.Title) < int(1) {
		return fmt.Errorf("field Title min_len rule failed, current value: %d", len(p.Title))
	}
	if len(p.Content) < int(1) {
		return fmt.Errorf("field Content min_len rule failed, current value: %d", len(p.Content))
	}
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	return nil
}
func (p *CreateBookResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *DeleteBookRequest) IsValid() error {
	if p.BookId <= int64(0) {
		return fmt.Errorf("field NoteId gt rule failed, current value: %v", p.BookId)
	}
	return nil
}
func (p *DeleteBookResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *UpdateBookRequest) IsValid() error {
	if p.BookId <= int64(0) {
		return fmt.Errorf("field NoteId gt rule failed, current value: %v", p.BookId)
	}
	return nil
}
func (p *UpdateBookResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *QueryBookRequest) IsValid() error {
	if p.UserId <= int64(0) {
		return fmt.Errorf("field UserId gt rule failed, current value: %v", p.UserId)
	}
	if p.Offset < int64(0) {
		return fmt.Errorf("field Offset ge rule failed, current value: %v", p.Offset)
	}
	if p.Limit < int64(0) {
		return fmt.Errorf("field Limit ge rule failed, current value: %v", p.Limit)
	}
	return nil
}
func (p *QueryBookResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *MGetBookRequest) IsValid() error {
	if len(p.BookIds) < int(1) {
		return fmt.Errorf("field NoteIds MinLen rule failed, current value: %v", p.BookIds)
	}
	return nil
}
func (p *MGetBookResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}