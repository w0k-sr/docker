<!DOCTYPE html>

<html>
<head>
  <title>問題</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="../static/css/modern-normalize.css">
  <link rel="stylesheet" type="text/css" href="../static/css/main.css">
  <link rel="stylesheet" type="text/css" href="../static/css/sudoku.css">
  <script src="../static/js/sudoku/detail.js"></script>
</head>

<body>
  <div class="min-h-screen bg-neutral-300 relative">
    <div class="pt-16 md:pt-20 pb-28 container px-2 md:px-5 xl:px-20">
      <div class="m-auto">
        <div class="mb-3 py-3 px-2 lg:px-10 rounded-t-md bg-neutral-600">
          <h1 class="text-neutral-100">{{.Sudoku.SudokuGenerate.Id}}</h1>
        </div>
        <div class="mb-3">
        </div>
        <div class="mb-3">
          <div class="center">
            <div class="grid grid-cols-9 sudoku-parent">
              {{range $i, $array := .preArray}}
                {{range $j, $v := $array}}
                  {{if eq $v 0}}
                    {{ $arrayLen := len $.solveArray }}
                    {{ $rowExists := lt $i $arrayLen }}
                    {{ $colExists := and $rowExists (lt $j (len (index $.solveArray $i))) }}
                    {{ $valueExists := and $colExists (ne (index $.solveArray $i $j) 0) }}

                    {{if $valueExists}}
                      <input class="bg-white sudoku-child scol_{{$j}} srow_{{$i}}" id="number" 
                      type="number" min="1" max="9" value="{{index $.solveArray $i $j}}">
                    {{else}}
                      <input class="bg-white sudoku-child scol_{{$j}} srow_{{$i}}" id="number" 
                      type="number" min="1" max="9" value="">
                    {{end}}
                  {{end}}
                  {{if ne $v 0}}
                    <div class="bg-neutral-600 sudoku-child scol_{{$j}} srow_{{$i}}">
                      {{$v}}
                    </div>
                  {{end}}
                {{end}}
              {{end}}
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</body>
</html>