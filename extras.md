/*resp, err := http.Get(url)
if err != nil {
  return nil, err
}
if resp.StatusCode != http.StatusOK {
  resp.Body.Close()
  return nil, fmt.Errorf("getting %s: %s", resp, resp.StatusCode)
}

doc, err := html.Parse(resp.Body)
resp.Body.Close()
if err != nil {
  return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
}*/

/*var nodesByID map[string]*html.Node

var f func(*html.Node)
f = func(n *html.Node) {
  if n.Type == html.ElementNode && n.Data == "h1" {
    for _, h := range n.Attr {
      if h.Key != "id" {
        continue
      }
      id := h.Val
      // make Amazon ID's Title field == value inside this element
      if id == "aiv-content-title" && n.FirstChild.Type == html.TextNode {
        amznID.Title = n.FirstChild.Data
      }
    }
  }
  for c := n.FirstChild; c != nil; c = c.NextSibling {
      f(c)
  }
}
f(doc)
*/
