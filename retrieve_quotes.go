/* Given a specific date and a list of stock symbols, this program returns
 * stock data for each stock.
 * Takes two command line parameters:
 * -date="Date"
 * -symbols="List symbols"
 * Example: retrieve_quotes.go -date="2016-10-17" -symbols="YHOO GOOG NFLX"
 */
package main
import (
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "strings"
    "time"
    "encoding/json"
    "flag"
)

// List of symbols in your portfolio. Take from stdin or file later. xxx/TODO
var symbols = []string{
    "NMBL",
    "VMW",
    "TWTR",
    "COP",
}

var (
    YahooTables = Tables{
        Quotes: "yahoo.finance.quotes",
        Hist: "yahoo.finance.historicaldata",
        DataTablesUrl: "store://datatables.org/alltableswithkeys",
        PublicAPIUrl: "http://query.yahooapis.com/v1/public/yql",
    }
)

// Struct with names of various tables in Yahoo's database
type Tables struct{
    Quotes string
    Hist string
    DataTablesUrl string
    PublicAPIUrl string
}

type JsonBase struct {
    count int
    created time.Time
    lang string
}

type IndividualStock struct {
    Symbol  string
    Name    string
    Date    string
    Open    string
    High    string
    Low     string
    Close   string
    Volume  string
    Adj_Close string
}

type JsonResponse struct {
    Query struct {
        JsonBase
        Results struct {
                    Quote []IndividualStock `json:"quote"`
         }
     }
}

type StockName struct {
    Name string
}

type JsonResponseName struct {
    Query struct {
        JsonBase
        Results struct {
            Quote StockName `json:"quote"`
                }
          }
}

// Given a query string as input, this functionr returns the raw bytes in the
// HTTP body obtained
func run_query(query string) ([]byte, error) {

    v := url.Values{}
    v.Set("q", query)
    v.Set("format", "json")
    v.Set("env", YahooTables.DataTablesUrl)

    url := YahooTables.PublicAPIUrl + "?" + v.Encode()
    //fmt.Println(url)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    httpBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    /* the first row includes column headers, ignore */
    return httpBody, nil
}

func MapStr(mapping func(string) string, xs []string) []string {
    mxs := make([]string, 0, len(xs))
    for _, s := range xs {
        mxs = append(mxs, mapping(s))
    }
    return mxs
}

// Given a stock symbol, retrieve the name of the company
func query_name(symbol string) (string, error) {
    query := fmt.Sprintf(`SELECT Name from %s where symbol = "%s"`,
                     YahooTables.Quotes, symbol)
    raw, err := run_query(query)
    if (err != nil) {
        return "", err
    }
    var resp JsonResponseName
    err = json.Unmarshal(raw, &resp)
    if (err != nil) {
        return "", err
    }
    //fmt.Println("query_name: ", resp)
    return resp.Query.Results.Quote.Name, nil
}

func query_data(date *string, symbols *string) error {
    start := fmt.Sprintf(` AND startDate = "%s"`, *date)
    end := fmt.Sprintf(` AND endDate = "%s"`, *date)

    syms := strings.Split(*symbols, " ")
    quotedSymbols := MapStr(func(s string) string {
        return `"` + s + `"`
    }, syms)
    query := fmt.Sprintf(`SELECT * FROM %s WHERE symbol IN (%s) %s %s`,
        YahooTables.Hist, strings.Join(quotedSymbols, ","), start, end)
    //fmt.Println(query)
    raw, err := run_query(query)
    if err != nil {
        return err
    }

    //fmt.Println(string(raw))
    var resp JsonResponse
    err = json.Unmarshal(raw, &resp)
    if (err != nil) {
        fmt.Println(err)
        return err
    }
    fmt.Println("Symbol  Name    Open  Close  Low   High  Volume  Adj_Close")
    for _, s := range resp.Query.Results.Quote {
        s.Name, err = query_name(s.Symbol)
        if err != nil {
            fmt.Println("Error obtaining company name for ", s.Symbol, ": ", err)
            return err
        }
        fmt.Println(s.Symbol, " ", s.Name, "   ", s.Open, " ", s.Close, " ", s.Low, " ", s.High, " ", s.Volume, " ", s.Adj_Close);
    }
    return nil
    //fmt.Println(resp)

}

func main() {
    datePtr := flag.String("date", "2016-10-20", "date for which data is needed")
    symbolsPtr := flag.String("symbols", "nmbl", "all stock symbols")
    flag.Parse()
    fmt.Println(*datePtr)
    fmt.Println(*symbolsPtr)

    err := query_data(datePtr, symbolsPtr)
    if (err != nil) {
        fmt.Println("query_data returned error: ", err)
        return
    }
}
