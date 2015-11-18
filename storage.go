package agro

import "github.com/barakmich/agro/models"

type Store interface {
	Flush() error
	Close() error
}

type INodeRef struct {
	Volume int64
	INode  int64
}

type BlockID struct {
	INodeRef
	Index int64
}

type BlockStore interface {
	Store
	GetBlock(b BlockID) ([]byte, error)
	WriteBlock(b BlockID, data []byte) error
	DeleteBlock(b BlockID) error
}

type INodeStore interface {
	Store
	GetINode(i uint64) (*models.INode, error)
	WriteINode(i uint64, inode *models.INode) error
	DeleteINode(i uint64) error
}
