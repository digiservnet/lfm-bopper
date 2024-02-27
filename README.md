# CSV -> JSON converter for LFM ACC BoP

LFM (Low Fuel Motorsport) run custom BoP data and render this on their BoP page. Copy & pasting that data into a spreadsheet allows for easy CSV creation. This tool will then convert that CSV file into a JSON file ready to apply to your own ACC server.

Usage:
```bash
lfm-bopper --in=<path/to/your_bop_file.csv>
```

Optionally, specify an output file (`bop.json` by default):
```bash
lfm-bopper --in=<path/to/your_bop_file.csv> --out=<path/to/your/bop.json>
```

Both GT3 and GT4 cars are currently handled.

`build` directory includes both Linux and Windows (AMD64) binaries.