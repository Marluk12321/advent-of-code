package mapping

import "fmt"

type State struct {
	Face         Face
	FacePosition Position
	Facing       FacingDirection
}

func (state State) String() string {
	return fmt.Sprint(
		"State(face: ", state.Face,
		", position on face: ", state.FacePosition, ", facing: ", state.Facing, ")")
}
