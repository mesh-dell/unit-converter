# Unit Converter Application

A small web application that provides quick unit conversions for **Length**, **Mass (Weight)**, and **Temperature**.
Backend logic is written in **Go** (uses maps for linear conversions and dedicated functions for temperature). Frontend uses Go `html/template` and is styled with **Tailwind CSS**.

---

## Features

* Length conversions (Metric ⇄ Imperial): Meter, Kilometer, Centimeter, Foot, Inch, Mile
* Mass (Weight) conversions (Metric ⇄ Imperial): Kilogram, Gram, Pound, Ounce
* Temperature conversions (non-linear): Celsius, Fahrenheit, Kelvin
* Minimal, responsive UI using Tailwind CSS

---

## Supported conversions (summary)

### Length

* Metric: `meter (m)`, `kilometer (km)`, `centimeter (cm)`
* Imperial: `foot (ft)`, `inch (in)`, `mile (mi)`

### Mass (Weight)

* Metric: `kilogram (kg)`, `gram (g)`
* Imperial: `pound (lb)`, `ounce (oz)`

### Temperature

* `Celsius (°C)` ⇄ `Fahrenheit (°F)` ⇄ `Kelvin (K)` (handled with dedicated functions because conversions are non-linear)

---

## Quick start

### Requirements

* Go 1.20+ (or your project's required Go version)
* Node / npm only needed if you rebuild Tailwind locally (optional). Pre-built CSS can be served from `static/` instead.

### Run locally

1. Ensure your Go source files (handlers, conversion logic, templates, static assets) are in the project directory.
2. Start the server:

```bash
go run main.go
```

3. Open the app in your browser:

```
http://localhost:8080
```

https://roadmap.sh/projects/unit-converter
