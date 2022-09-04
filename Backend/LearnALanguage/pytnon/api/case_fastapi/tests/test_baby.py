import pytest

from baby_monstar import baby


def test_dummy():
    assert "chome" == "chome"


class TestBaby:
    @pytest.fixture(scope="class")
    def generate_test_baby(self):
        test_baby = baby.Baby(
            name="bebi",
            birth_month=1,
            phys_sex="men",
        )
        return test_baby

    def test_baby_cry_ok(self, generate_test_baby):
        expect = "Gyaaaaa!!!"
        actual = generate_test_baby.cry()
        assert expect == actual

    def test_baby_poop_ok(self, generate_test_baby):
        expect = "💩💩💩"
        actual = generate_test_baby.poop()
        assert expect == actual

    def test_baby_milk_ok(self, generate_test_baby):
        expect = "🍼🍼🍼"
        actual = generate_test_baby.milk()
        assert expect == actual
