from app.monster import *

def test_roar_ok():
  data = Monster(
      name = "Araminasu",
      hitpoint = 10,
      atack = 30,
    )
  expect = "Shaaah"
  result = data.roar

  assert expect == result

def test_desc_ok():
  data = Monster(
      name = "Ikeremil",
      hitpoint = 15,
      atack = 20,
    )
  expect = "name = Ikeremil hitpoint = 15 atack = 20 roar = Shaaah"
  result = data.decribe()

  assert expect == result
