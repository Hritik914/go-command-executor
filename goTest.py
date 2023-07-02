import requests

def run_test(label, response, expected_status, expected_output=None):
    try:
        status_code = response.status_code
        if status_code == expected_status:
            if expected_output is not None:
                output = response.json().get("output", "")
                if output == expected_output:
                    print(f"{label} test passed")
                else:
                    print(f"{label} test failed: Unexpected output")
            else:
                print(f"{label} test passed")
        else:
            print(f"{label} test failed: Unexpected status code")
    except AttributeError:
        print(f"{label} test failed: Error occurred during the request")

def test_get_command():
    url = "http://localhost:8080/api/cmd"
    params = {
        "command": "echo hello"
    }
    response = requests.get(url, params=params)
    expected_output = "hello\n"
    run_test("GET command", response, 200, expected_output)

def test_post_command():
    url = "http://localhost:8080/api/cmd"
    payload = {
        "command": "echo test"
    }
    response = requests.post(url, json=payload)
    expected_output = "test\n"
    run_test("POST command", response, 200, expected_output)

def test_missing_get_parameter():
    url = "http://localhost:8080/api/cmd"
    response = requests.get(url)
    run_test("Missing GET parameter", response, 400)

def test_missing_post_body():
    url = "http://localhost:8080/api/cmd"
    response = requests.post(url)
    run_test("Missing POST body", response, 400)

def test_invalid_get_command():
    url = "http://localhost:8080/api/cmd"
    params = {
        "command": "invalidcommand"
    }
    response = requests.get(url, params=params)
    run_test("Invalid GET command", response, 404)

def test_invalid_post_command():
    url = "http://localhost:8080/api/cmd"
    payload = {
        "command": "invalidcommand"
    }
    response = requests.post(url, json=payload)
    run_test("Invalid POST command", response, 404)

def run_tests():
    test_get_command()
    test_post_command()
    test_missing_get_parameter()
    test_missing_post_body()
    test_invalid_get_command()
    test_invalid_post_command()

if __name__ == "__main__":
    run_tests()
