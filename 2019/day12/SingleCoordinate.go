package main

type System1D struct {
	pos []int64
	vel []int64
}

func NewSystem1D(pos []int64) *System1D {
	system := &System1D{
		pos: make([]int64, len(pos)),
		vel: make([]int64, len(pos)),
	}
	copy(system.pos, pos)
	return system
}

func (s *System1D) RunUntilCycle() (nbSteps int) {
	initialSystem := NewSystem1D(s.pos)

	for {
		nbSteps++
		s.Update()
		if s.Equal(initialSystem) {
			break
		}
	}
	return nbSteps
}

func (s *System1D) Update() {
	size := len(s.pos)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			veli, velj := compare2(s.pos[i], s.pos[j])
			s.vel[i] += veli
			s.vel[j] += velj
		}
		s.pos[i] += s.vel[i]
	}
}

func (s *System1D) Equal(other *System1D) bool {
	for i := range s.pos {
		if s.pos[i] != other.pos[i] || s.vel[i] != other.vel[i] {
			return false
		}
	}
	return true
}

func compare2(i, j int64) (int64, int64) {
	if i < j {
		return 1, -1
	}
	if i > j {
		return -1, 1
	}
	return 0, 0
}
