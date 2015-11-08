package patcher_test

import (
	"testing"
	"time"

	p "github.com/lysu/go-struct-patcher"
	"github.com/stretchr/testify/assert"
)

type Comment struct {
	NickName string
	Content  string
	Date     time.Time
}

type Blog struct {
	Title      string
	CommentIds []uint64
	Comments   map[string]*Comment
}

func (b Blog) FirstComment() *Comment {
	return b.Comments["0"]
}

func TestDoPatch(t *testing.T) {
	assert := assert.New(t)
	patcher := p.Patcher{}
	b := &Blog{
		Title:      "Blog title1",
		CommentIds: []uint64{1, 3},
		Comments: map[string]*Comment{
			"0": {
				NickName: "000",
				Content:  "test",
				Date:     time.Now(),
			},
			"1": {
				NickName: "u1",
				Content:  "test",
				Date:     time.Now(),
			},
			"3": {
				NickName: "tester",
				Content:  "test hehe...",
				Date:     time.Now(),
			},
		},
	}
	ps := p.Patch{
		"title":                            "title B",
		"commentIds[1]":                    uint64(4),
		"firstComment().content":           "hehe~",
		"comments[commentIds[0]].nickName": "私",
	}
	err := patcher.PatchIt(b, ps)
	assert.NoError(err)
	assert.Equal("title B", b.Title)
	assert.Equal(uint64(4), b.CommentIds[1])
	assert.Equal("私", b.Comments["1"].NickName)
	assert.Equal("hehe~", b.FirstComment().Content)
}
