package main

/*func TestAdd(t *testing.T) {
	assert.Equal(t, int64(10), Add(6, 4))
	assert.Equal(t, int64(0), Add(0, 0))
	assert.Equal(t, int64(5), Add(10, -5))
}

func TestMul(t *testing.T) {
	assert.Equal(t, int64(6), Multiply(3, 2))
	assert.Equal(t, int64(0), Multiply(5, 0))
	assert.Equal(t, int64(-4), Multiply(2, -2))
	assert.Equal(t, int64(1), Multiply(-1, -1))
}

func TestIn(t *testing.T) {
	in := "5"
	out := new(strings.Builder)
	assert.Equal(t, int64(5), Input(out, strings.NewReader(in)))
	assert.NotEmpty(t, out.String())

	in = "abcd"
	out = new(strings.Builder)
	assert.Panics(t, func() { Input(out, strings.NewReader(in)) })
}

func TestOut(t *testing.T) {
	out := new(strings.Builder)
	Output(out, 5)
	assert.Contains(t, out.String(), "5")
}

func TestJNZ(t *testing.T) {
	assert.Equal(t, 1337, JumpNonZero(1, 1337, 42))
	assert.Equal(t, 0, JumpNonZero(42, 0, 42))
	assert.Equal(t, 42, JumpNonZero(0, 1337, 42))
}

func TestJZ(t *testing.T) {
	assert.Equal(t, 1337, JumpZero(0, 1337, 42))
	assert.Equal(t, 0, JumpZero(42, 1337, 0))
	assert.Equal(t, 42, JumpZero(1, 1337, 42))
}

func TestLT(t *testing.T) {
	assert.Equal(t, int64(1), LessThan(42, 1337))
	assert.Equal(t, int64(0), LessThan(1337, 42))
}

func TestEq(t *testing.T) {
	assert.Equal(t, int64(1), Equal(42, 42))
	assert.Equal(t, int64(0), Equal(1337, 42))
}

func TestRelBase(t *testing.T) {
	assert.Equal(t, int64(15), ChangeRelativeBase(5, 10))
	assert.Equal(t, int64(2), ChangeRelativeBase(-3, 5))
	assert.Equal(t, int64(-2), ChangeRelativeBase(3, -5))
	assert.Equal(t, int64(-8), ChangeRelativeBase(-3, -5))
}

func TestBoolToInt(t *testing.T) {
	assert.Equal(t, int64(1), boolToInt(true))
	assert.Equal(t, int64(0), boolToInt(false))
}

func TestIntToBool(t *testing.T) {
	assert.Equal(t, true, intToBool(5))
	assert.Equal(t, true, intToBool(-3))
	assert.Equal(t, false, intToBool(0))
}*/
