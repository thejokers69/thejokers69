package handlers

import (
    "net/http/httptest"
    "os"
    "path/filepath"
    "testing"

    "github.com/gin-gonic/gin"
)

func TestStaticHandler_AllowedFile(t *testing.T) {
    tmp, err := os.MkdirTemp("", "static-test-allowed")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(tmp)

    // Create static/allowed/index.txt
    staticRoot := filepath.Join(tmp, "static")
    if err := os.MkdirAll(filepath.Join(staticRoot, "allowed"), 0o755); err != nil {
        t.Fatal(err)
    }
    filePath := filepath.Join(staticRoot, "allowed", "index.txt")
    if err := os.WriteFile(filePath, []byte("ok"), 0o644); err != nil {
        t.Fatal(err)
    }

    // chdir to tmp
    orig, _ := os.Getwd()
    _ = os.Chdir(tmp)
    defer os.Chdir(orig)

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    req := httptest.NewRequest("GET", "/static/allowed/index.txt", nil)
    c.Request = req
    c.Params = gin.Params{gin.Param{Key: "filepath", Value: "static/allowed/index.txt"}}
    StaticHandler(c)
    if w.Code != 200 {
        t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
    }
}

func TestStaticHandler_PathTraversalBlocked(t *testing.T) {
    tmp, err := os.MkdirTemp("", "static-test-traversal")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(tmp)
    staticRoot := filepath.Join(tmp, "static")
    _ = os.MkdirAll(staticRoot, 0o755)

    orig, _ := os.Getwd()
    _ = os.Chdir(tmp)
    defer os.Chdir(orig)

    // Attempt to access outside via traversal
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    req := httptest.NewRequest("GET", "/static/../etc/passwd", nil)
    c.Request = req
    c.Params = gin.Params{gin.Param{Key: "filepath", Value: "static/../etc/passwd"}}
    StaticHandler(c)
    if w.Code != 400 {
        t.Fatalf("expected 400 for traversal, got %d", w.Code)
    }
}

func TestStaticHandler_EncodedTraversalBlocked(t *testing.T) {
    tmp, err := os.MkdirTemp("", "static-test-encoded")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(tmp)
    staticRoot := filepath.Join(tmp, "static")
    _ = os.MkdirAll(staticRoot, 0o755)

    orig, _ := os.Getwd()
    _ = os.Chdir(tmp)
    defer os.Chdir(orig)

    // Encoded traversal attempt
    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    req := httptest.NewRequest("GET", "/static/%2e%2e/etc/passwd", nil)
    c.Request = req
    c.Params = gin.Params{gin.Param{Key: "filepath", Value: "static/%2e%2e/etc/passwd"}}
    StaticHandler(c)
    if w.Code != 400 {
        t.Fatalf("expected 400 for encoded traversal, got %d", w.Code)
    }
}

func TestStaticHandler_SymlinkEscapeBlocked(t *testing.T) {
    tmp, err := os.MkdirTemp("", "static-test-symlink")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(tmp)

    staticRoot := filepath.Join(tmp, "static")
    if err := os.MkdirAll(staticRoot, 0o755); err != nil {
        t.Fatal(err)
    }
    // outside secret
    outside := filepath.Join(tmp, "outside")
    if err := os.MkdirAll(outside, 0o755); err != nil {
        t.Fatal(err)
    }
    secretFile := filepath.Join(outside, "secret.txt")
    if err := os.WriteFile(secretFile, []byte("secret"), 0o644); err != nil {
        t.Fatal(err)
    }

    // create a symlink inside static pointing outside
    linkPath := filepath.Join(staticRoot, "link")
    if err := os.Symlink(secretFile, linkPath); err != nil {
        t.Fatal(err)
    }

    orig, _ := os.Getwd()
    _ = os.Chdir(tmp)
    defer os.Chdir(orig)

    gin.SetMode(gin.TestMode)
    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)
    req := httptest.NewRequest("GET", "/static/link", nil)
    c.Request = req
    c.Params = gin.Params{gin.Param{Key: "filepath", Value: "static/link"}}
    StaticHandler(c)
    if w.Code != 400 {
        t.Fatalf("expected 400 for symlink escape, got %d", w.Code)
    }
}
