# ⚡ CGO Configuration (`CGO_ENABLED=0`)

This project explicitly disables CGO (`CGO_ENABLED=0`) across development, testing, and production build pipelines.

---

## 🎯 Rationale

Go's default behavior enables CGO when a local C compiler (such as `gcc` or `clang`) is detected. However, disabling CGO and building pure Go binaries aligns with industry best practices for Go microservices and applications.

### Key Benefits

- **Purely Static Binaries:** Disabling CGO ensures the resulting binary has zero dynamic C library dependencies (such as `glibc` or `musl`). This allows the binary to run inside minimal Docker containers (e.g., `scratch` or `alpine`) without runtime crashes related to missing C libraries.
- **Faster Local Reloads:** Live-reloading tools like `air` compile significantly faster because the Go toolchain (`gc`) doesn't need to invoke external C compilers on every file change.
- **Seamless Cross-Compilation:** Cross-compiling for different architectures or target operating systems (e.g., compiling for `linux/arm64` from an `x86_64` workstation) works out of the box without requiring complex C cross-compilation toolchains.
- **Environment Independence:** Guarantees consistent builds across isolated environments, non-standard distros (such as **NixOS**), CI/CD runners, and developer machines without requiring system-level C toolchains.

---

## 🛠️ Local Development Usage

When running the application locally or using hot-reloading tools, ensure `CGO_ENABLED=0` is passed:

```bash
# Direct run
CGO_ENABLED=0 go run main.go

# Direct build
CGO_ENABLED=0 go build -o bin/app .

# Live-reload with Air
CGO_ENABLED=0 air
```
