<!DOCTYPE html>

<html>
<head>
  <title>問題</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Hachi+Maru+Pop&display=swap" rel="stylesheet">
  <link rel="stylesheet" href="../static/css/modern-normalize.css">
  <link rel="stylesheet" type="text/css" href="../static/css/main.css">
  <link rel="stylesheet" type="text/css" href="../static/css/sudoku.css">
  <script src="https://code.jquery.com/jquery-3.7.1.min.js" 
  integrity="sha256-/JqT3SQfawRcv/BIHPThkBvs0OEvtFFmqPF/lYI/Cxo=" crossorigin="anonymous"></script>
  <script src="../static/js/sudoku/detail.js"></script>
</head>

<body>
  <div class="min-h-screen bg-neutral-300 relative">
    <div class="pt-16 md:pt-20 pb-28 container px-2 md:px-5 xl:px-20">
      <div class="m-auto">
        <div class="py-2 px-2 lg:px-10 rounded-t-md bg-neutral-600">
          <h1 class="text-neutral-100">No.{{.Sudoku.SudokuGenerate.Id}}</h1>
        </div>
        <div class="mb-3">
        </div>
        <div class="mb-3">
          <div class="center flex">
            <!-- 数独 -->
            <div class="grid grid-cols-9 sudoku-parent">
              {{range $i, $array := .preArray}}
                {{range $j, $v := $array}}
                  {{if eq $v 0}}
                    {{ $arrayLen := len $.solveArray }}
                    {{ $rowExists := lt $i $arrayLen }}
                    {{ $colExists := and $rowExists (lt $j (len (index $.solveArray $i))) }}
                    {{ $valueExists := and $colExists (ne (index $.solveArray $i $j) 0) }}

                    {{if $valueExists}}
                      <input class="bg-white sudoku-child" id="number" type="number" 
                      min=1 max=9 value="{{index $.solveArray $i $j}}" data-col={{$j}} data-row={{$i}} data-check="{{index $.flagArray $i $j}}">
                    {{else}}
                      <input class="bg-white sudoku-child" id="number" type="number" 
                      min=1 max=9 value="" data-col={{$j}} data-row={{$i}} data-check="{{index $.flagArray $i $j}}">
                    {{end}}
                  {{end}}
                  {{if ne $v 0}}
                    <div class="bg-gray-300 sudoku-child" data-col={{$j}} data-row={{$i}} data-check="{{index $.flagArray $i $j}}">
                      {{$v}}
                    </div>
                  {{end}}
                {{end}}
              {{end}}
            </div>
            <!-- 画像 -->
            <div class="neko" data-neko="{{if eq .comment 1}}false{{else}}true{{end}}">
              <div class="comment">
                <p class="comment-run">まだまだ<br>走るよ～♪</p>
              </div>
              <div class="neko-run"></div>
            </div>
            <div class="neko" data-neko="{{if eq .comment 2}}false{{else}}true{{end}}">
              <div class="comment">
                <p class="comment-cry">おっと・・・<br>どれかわかるかな</p>
              </div>
              <div class="neko-cry"></div>
            </div>
            <div class="neko" data-neko="{{if eq .comment 3}}false{{else}}true{{end}}">
              <div class="comment">
                <p class="comment-smile">この調子だ♪<br>もうひと踏ん張り</p>
              </div>
              <div class="neko-smile"></div>
            </div>
            <div class="neko" data-neko="{{if eq .comment 4}}false{{else}}true{{end}}">
              <div class="comment">
                <p class="comment-finish">完成♪<br>よくやったね～</p>
              </div>
              <div class="neko-finish"></div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <input id="sudoku_id" type="hidden" value={{.Sudoku.SudokuGenerate.Id}} />
    <input id="comment_num" type="hidden" value={{.comment}} />
  </div>
</body>
</html>