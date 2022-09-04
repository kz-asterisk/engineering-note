import pytest
from fastapi.testclient import TestClient

from baby_monstar import baby
from main import app

client = TestClient(app)


class TestMainGet:
    @pytest.fixture(scope="class")
    def generate_test_baby(self):
        test_baby = baby.Baby(
            name="bebibebi",
            birth_month=2,
            pys_sex="women",
        )
        return test_baby

    def test_get_root_ok(self, generate_test_baby):
        _ = generate_test_baby
        response = client.get("/")
        assert response.status_code == 200
        assert response.json() == {"message": "ここには何もないよ"}

    def test_get_hello_ok(self, generate_test_baby):
        _ = generate_test_baby
        response = client.get("/hello")
        assert response.status_code == 200
        assert response.json() == {"message": "Hi! I'm baby monser👶👶👶"}

    def test_get_cry_ok(self, generate_test_baby):
        _ = generate_test_baby
        response = client.get("/cry")
        assert response.status_code == 200
        assert response.json() == {"message": "Gyaaaaa!!!"}

    def test_get_poop_ok(self, generate_test_baby):
        _ = generate_test_baby
        response = client.get("/poop")
        assert response.status_code == 200
        assert response.json() == {"message": "💩💩💩"}

    def test_get_milk_ok(self, generate_test_baby):
        _ = generate_test_baby
        response = client.get("/milk")
        assert response.status_code == 200
        assert response.json() == {"message": "🍼🍼🍼"}


class TestMainPut:
    @pytest.fixture(scope="class")
    def generate_test_baby(self):
        test_baby = baby.Baby(
            name="bebibebi",
            birth_month=2,
            pys_sex="women",
        )
        return test_baby
