// curl -siL https://courageous-cupcake-6d9019.netlify.app/.netlify/functions/start
// https://hub.fastgit.xyz/aws/aws-lambda-go/blob/main/events/apigw.go

package main

import (
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "io/ioutil"
  "net/http"
  "log"
  "time"
  "fmt"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
  log.Println("path:", request.PathParameters)
  url := request.QueryStringParameters["url"]
  // body := fetch(url)
  // log.Println("body:", body)
  body := add(100, 200)
  log.Println("body:", body)
  return &events.APIGatewayProxyResponse{
    StatusCode:        200,
    Headers:           map[string]string{"Content-Type": "text/plain"},
    MultiValueHeaders: http.Header{"Set-Cookie": {"Ding", "Ping"}},
    Body:              fmt.Sprintf("> URL: %s %s", url, body);
    // Body:              fmt.Sprintf("> Body: %s", body),
    IsBase64Encoded:   false,
  }, nil
}

func add(m, n int) int {
  return m+n
}

func fetch(url string) string {
  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  if resp.StatusCode != 200 {
    log.Fatalf("Failed to fetch data: %d %s", resp.StatusCode, resp.Status)
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }
  // []byte -> string
  strBody := string(body)
  // log.Print(strBody)
  return strBody
  // doc, err := goquery.NewDocumentFromReader(resp.Body)
  // if err != nil {
  //   log.Fatal(err)
  // }
  // title := doc.Find("title").Text()
  // log.Info(title)
}

func main() {
  // url := "https://search.bilibili.com"
  // fetch(url)
  // log.Print(time.Second)
  // log.Print(fmt.Sprintf("%s DONE", "Task one"))

  // Make the handler available for Remote Procedure Call
  lambda.Start(handler)
}

