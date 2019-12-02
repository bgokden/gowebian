package mdb

import "github.com/bgokden/gowebian/pkg/component"

type Header struct {
	component.BaseComponent
}

func NewHeader() *Header {
	h := &Header{}
	h.SetTag("header")
	return h
}

type Main struct {
	component.BaseComponent
}

func NewMain() *Main {
	m := &Main{}
	m.SetTag("main")
	return m
}

type Container struct {
	component.BaseComponent
}

func NewContainer() *Container {
	c := &Container{}
	c.SetAttribute("class", "container")
	return c
}

type Row struct {
	component.BaseComponent
}

func NewRow() *Row {
	c := &Row{}
	c.SetAttribute("class", "row")
	return c
}

type Col struct {
	component.BaseComponent
}

func NewCol() *Col {
	c := &Col{}
	c.SetAttribute("class", "col")
	return c
}
