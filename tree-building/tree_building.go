package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID int
	Parent int
}

type Node struct {
	ID int
	Children []*Node
}

type ById []Record

func (a ById) Len() int 			{ return len(a) }
func (a ById) Swap(i, j int) 		{ a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool	{ return a[i].ID < a[j].ID }

func Build(records []Record) (*Node, error) {
	numberOfRoots := 0
	nodes := map[int]*Node{0: {ID: 0}}

	if len(records) == 0 {
		return nil, nil
	}

	sort.Sort(ById(records))

	for _, record := range records {
		if _, ok := nodes[record.ID]; ok && record.ID != 0  {
			return nil, errors.New("duplicate node")
		}

		if record.ID < 0 || record.ID > len(records)-1 || (record.ID != 0 && record.ID <= record.Parent) || (record.ID == 0 && record.Parent != 0) {
			return nil, errors.New("bad records")
		}

		if record.ID == 0 {
			numberOfRoots ++
		} else {
			nodes[record.ID] = &Node{ID: record.ID}
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, nodes[record.ID])
		}

		if numberOfRoots > 1 {
			return nil, errors.New("multiple roots")
		}
	}


	root, _ := nodes[0]
	return root, nil
}