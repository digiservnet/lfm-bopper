package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

type BopRecord struct {
    Track    string `json:"track"`
    CarModel int    `json:"carModel"`
    Ballast  int    `json:"ballastKg"`
}

type BopEntries struct {
    Entries []BopRecord `json:"entries"`
}

func createBopRecords(data [][]string) []BopRecord {
    // Convert CSV to array of BopRecords
    var bopRecords []BopRecord

    if len(data) < 1 {
        fmt.Println("Invalid BoP data file.")
        os.Exit(1)
    }

    for _, line := range data {
        var bopRecord BopRecord
        var carName string
        var carYear string

        for lines, field := range line {
            switch lines {
            case 1:
                carName = field
            case 2:
                carYear = field
            case 3:
                bopRecord.Ballast, _ = strconv.Atoi(strings.Replace(field, " kg", "", -1))
            case 5:
                bopRecord.Track = field
            }
        }

        bopRecord.CarModel = carModelFromName(carName, carYear)

        bopRecords = append(bopRecords, bopRecord)
    }

    return bopRecords
}

func carModelFromName(name string, year string) int {
    var carModel int

    car := strings.Join([]string{name, year}, " ")

    switch car {
    case "Porsche 991 GT3 R 2018":
        carModel = 0
    case "Mercedes-AMG GT3 2015":
        carModel = 1
    case "Ferrari 488 GT3 2018":
        carModel = 2
    case "Audi R8 LMS 2015":
        carModel = 3
    case "Lamborghini Huracan GT3 2015":
        carModel = 4
    case "McLaren 650S GT3 2015":
        carModel = 5
    case "Nissan GT-R Nismo GT3 2018":
        carModel = 6
    case "BMW M6 GT3 2017":
        carModel = 7
    case "Bentley Continental GT3 2018":
        carModel = 8
    // LFM-specific identifier
    case "Bentley Continental 2018":
        carModel = 8
    case "Porsche 991 II GT3 Cup 2017":
        carModel = 9
    case "Nissan GT-R Nismo GT3 2015":
        carModel = 10
    case "Bentley Continental GT3 2015":
        carModel = 11
    case "AMR V12 Vantage GT3 2013":
        carModel = 12
    case "Reiter Engineering R-EX GT3 2017":
        carModel = 13
    case "Emil Frey Jaguar G3 2012":
        carModel = 14
    case "Lexus RC F GT3 2016":
        carModel = 15
    case "Lamborghini Huracan GT3 Evo 2019":
        carModel = 16
    case "Honda NSX GT3 2017":
        carModel = 17
    case "Lamborghini Huracan SuperTrofeo 2015":
        carModel = 18
    case "Audi R8 LMS Evo 2019":
        carModel = 19
    case "AMR V8 Vantage 2019":
        carModel = 20
    case "Honda NSX GT3 Evo 2019":
        carModel = 21
    case "McLaren 720S GT3 2019":
        carModel = 22
    case "Porsche 991 II GT3 R 2019":
        carModel = 23
    // LFM-specific identifier
    case "Porsche 991II GT3 R 2019":
        carModel = 23
    case "Ferrari 488 GT3 Evo 2020":
        carModel = 24
    case "Mercedes-AMG GT3 2020":
        carModel = 25
    case "BMW M4 GT3 2021":
        carModel = 26
    case "BMW M2 Club Sport Racing 2020":
        carModel = 27
    case "Porsche 992 GT3 Cup 2021":
        carModel = 28
    case "Lamborghini Huracan SuperTrofeo EVO2 2021":
        carModel = 29
    case "Ferrari 488 Challenge Evo 2020":
        carModel = 30
    case "Audi R8 LMS GT3 evo II 2022":
        carModel = 31
    case "Ferrari 296 GT3 2023":
        carModel = 32
    case "Lamborghini Huracan GT3 EVO 2 2023":
        carModel = 33
    case "Porsche 992 GT3 R 2023":
        carModel = 34
    case "McLaren 720S GT3 Evo 2023":
        carModel = 35
    case "Alpine A110 GT4 2018":
        carModel = 50
    case "Aston Martin Vantage GT4 2018":
        carModel = 51
    case "Audi R8 LMS GT4 2018":
        carModel = 52
    case "BMW M4 GT4 2018":
        carModel = 53
    case "Chevrolet Camaro GT4 2017":
        carModel = 55
    case "Ginetta G55 GT4 2012":
        carModel = 56
    case "KTM X-Bow GT4 2016":
        carModel = 57
    case "Maserati MC GT4 2016":
        carModel = 58
    case "McLaren 570S GT4 2016":
        carModel = 59
    case "Mercedes AMG GT4 2016":
        carModel = 60
    case "Porsche 718 Cayman GT4 Clubsport 2019":
        carModel = 61
    }

    return carModel
}

func main() {
    appName := os.Args[0]

    if len(os.Args[1:]) < 1 {
        fmt.Println("Please provide a BoP CSV file.")
        fmt.Println("")
        fmt.Println("Usage: " + appName + " <csv_file>")
        os.Exit(1)
    }

    bopCsvFile := os.Args[1]

    // Open CSV file
    file, err := os.Open(bopCsvFile)

    if err != nil {
        log.Fatal(err)
    }

    // Close CSV file at end of operation
    defer file.Close()

    // Read CSV file
    csvReader := csv.NewReader(file)
    data, err := csvReader.ReadAll()

    if err != nil {
        log.Fatal(err)
    }

    // Create BopRecords data
    bopRecords := createBopRecords(data)

    var bopEntries BopEntries
    bopEntries.Entries = bopRecords

    // Convert array of structs into JSON
    jsonData, err := json.MarshalIndent(bopEntries, "", "    ")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(jsonData))
}
