1. goes to chatgpt and use this prompt

Convert the following Go function into a D2 flowchart diagram.  
- Use explicit attribute‐object syntax (`node: { shape: ..., label: ... }`).  
- Use `shape: rectangle` for regular process steps, `shape: diamond` for decisions.  
- For any external or utility calls (e.g. `utils.*`, `fmt.*`, `time.*`, custom subfunctions), add  
style: {
stroke-dash: 3
}

scss
Copy
Edit
inside that node to render it with a dashed border.  
- Write human‐readable labels in Bahasa Indonesia.  
- Output **only** the valid D2 code block.

```go
// Paste your Go function here
func MyFunction(...) {
  // …
}

2. replace the function placeholder with the one that you need
3. copy the result and compile to https://play.d2lang.com/
4. replace the code into corresponding folder
5. feel free to add or delete the file as you need
6. push into a repo to get it regenerate for all