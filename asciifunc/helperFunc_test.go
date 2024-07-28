package asciifunc

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestAsciiConverter(t *testing.T) {
	result1 := "\n" + AsciiConverter("hello", "standard.txt")
	expected1 := "\n" + " _              _   _          \n" +
		"| |            | | | |         \n" +
		"| |__     ___  | | | |   ___   \n" +
		"|  _ \\   / _ \\ | | | |  / _ \\  \n" +
		"| | | | |  __/ | | | | | (_) | \n" +
		"|_| |_|  \\___| |_| |_|  \\___/  \n" +
		"                               \n" +
		"                               \n"

	if result1 != expected1 {
		t.Errorf("expected this  %v  \nbut got this  %v , the length of result is \n %v \n while length of expected is \n %v", expected1, result1, len(result1), len(expected1))
	}
}

func TestInputStringContainsNonASCIICharacters(t *testing.T) {
	input := "HelloWorld123©"
	result := IsItAnAsciiCharacter(input)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestConverts2DSliceToString(t *testing.T) {
	input := [][]string{
		{"Hello", "World"},
		{"Foo", "Bar"},
	}
	expected := "HelloWorld\nFooBar"

	result := Change2DtoString(input)

	if result != expected {
		t.Errorf("expected %s but got %s", expected, result)
	}
}

func TestHandleEmptyFileName(t *testing.T) {
	// Call the function under test with an empty file name
	result := CheckFileAuthencity("")

	// Assert the result
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

// Render template successfully with valid data and template path
func TestRenderTemplateSuccessfully(t *testing.T) {
	tmpl := "test_template.html"
	data := map[string]string{"Title": "Test Page"}
	statusCode := http.StatusOK

	// Create a temporary template file
	file, err := os.Create(tmpl)
	if err != nil {
		t.Fatalf("Failed to create template file: %v", err)
	}
	defer os.Remove(tmpl)
	defer file.Close()

	// Write some content to the template file
	_, err = file.WriteString("<html><head><title>{{.Title}}</title></head><body></body></html>")
	if err != nil {
		t.Fatalf("Failed to write to template file: %v", err)
	}

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		PageNotFound(w, tmpl, data, statusCode)
	})

	// Create a new request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Serve the request
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != statusCode {
		t.Errorf("handler returned wrong status code: got %v want %v", status, statusCode)
	}

	// Check the response body
	expected := "<html><head><title>Test Page</title></head><body></body></html>"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestStatusInternalServerError_ParsesTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	StatusInternalServerError(w, r)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestStatusUnavailableBanner_Success(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	StatusUnavailableBanner(w, r)

	result := w.Result()
	defer result.Body.Close()

	if result.StatusCode != 302 {
		t.Errorf("expected status OK; got %v", result.Status)
	}
}

// Successfully parses and executes the "400.html" template
func TestBadRequestExecutesTemplate(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/badrequest", nil)

	// Create a valid 400.html template file
	tmplContent := "<html><body>Bad Request</body></html>"
	ioutil.WriteFile("400.html", []byte(tmplContent), 0o644)
	defer os.Remove("400.html")

	BadRequest(w, r)

	result := w.Result()
	body, _ := ioutil.ReadAll(result.Body)

	if result.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", result.StatusCode)
	}

	if !strings.Contains(string(body), "Bad Request") {
		t.Errorf("Expected body to contain 'Bad Request', got %s", string(body))
	}
}

// Correctly parses the "405.html" template file

func TestCorrectTemplateParsing(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Wrongmethodused)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
