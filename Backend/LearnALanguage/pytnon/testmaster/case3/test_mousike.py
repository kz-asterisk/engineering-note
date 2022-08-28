import mousike

def test_tone_ok():
  data = mousike.Guitar(
    tone = "Chuweeen"
  )
  expect = "Chuweeen♪♪♪"
  result = data.sounds()

  assert expect == result

