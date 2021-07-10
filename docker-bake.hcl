variable "PKGVER" {
    default = "0.1.0"
}

group "default" {
    targets = ["deb"]
}

target "deb" {
    dockerfile = "Dockerfile"
    target = "final"
    output = ["artifacts"]
    args = {
        VERSION="${PKGVER}"
    }
}
