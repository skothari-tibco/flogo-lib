package model

import (
	"github.com/TIBCOSoftware/flogo-lib/core/ext/activity"
	"github.com/TIBCOSoftware/flogo-lib/core/process"
)

// ProcessContext is the execution context of the Process when executing
// a Process Behavior fuction
type ProcessContext interface {

	// ProcessDefinition returns the Process definition associated with this context
	ProcessDefinition() *process.Definition

	//State gets the state of the Process instance
	State() int

	//SetState sets the state of the Process instance
	SetState(state int)
}

// TaskContext is the execution context of the Task when executing
// a Task Behavior fuction
type TaskContext interface {

	// State gets the state of the Task instance
	State() int

	// SetState sets the state of the Task instance
	SetState(state int)

	// Task returns the Task associated with this context
	Task() *process.Task

	// FromLinks returns the set of predecessor Links of the current
	// task.
	FromLinks() []LinkContext

	// EnterLeadingChildren enters the set of child Tasks that
	// do not have any incoming links.
	// todo: should we allow cross-boundary links?
	//EnterLeadingChildren(int enterCode)

	// EnterChildren enters the set of child Tasks specified,
	// If single TaskEntry with nil Task is supplied,
	// all the child tasks are entered with the specified code.
	EnterChildren(taskEntries []*TaskEntry)

	// EvalLink evalutes the specified link, returning the resulting
	// LinkContext
	EvalLink(link *process.Link, code int) LinkContext

	//EvalLink(link *Link, code int) bool

	// Activity gets the Activity associated with the Task
	Activity() (activity activity.Activity, activityContext activity.Context)
}

// LinkContext is the execution context of the Task when executing
// a Link
// todo remove
type LinkContext interface {

	// Link returns the Link associated with this context
	Link() *process.Link

	// State gets the state of the Link instance
	State() int

	// SetState sets the state of the Link instance
	SetState(state int)
}