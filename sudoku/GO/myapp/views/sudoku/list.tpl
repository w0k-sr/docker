<!DOCTYPE html>

<html>
<head>
  <title>数独一覧</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" href="../static/css/modern-normalize.css">
  <link rel="stylesheet" type="text/css" href="../static/css/main.css">
  <script src="../static/js/sudoku/list.js"></script>
</head>

<body>
  <div class="min-h-screen bg-neutral-300 relative">
    <div class="pt-16 md:pt-20 pb-28 container px-2 md:px-5 xl:px-20">
      <div class="m-auto">
        <div class="mb-3 py-3 px-2 lg:px-10 rounded-t-md bg-neutral-600 flex flex-wrap flex-space">
          <h1 class="text-neutral-100">数独一覧</h1>
          <div class="mb-3 flex">
            <div class="items-center">
              {{.Message}}
            </div>
            <a href="/sudoku/create" class="btn btn-custom01">
              <span class="btn-custom01-front">問題を作成</span>
            </a>
          </div>
        </div>
        <div class="mb-3">
          <!-- <form id="comment" action="/sudoku/create/" method="post">
            <button type="submit" class="button">問題を作成</button>
          </form> -->
        </div>
        <div class="mb-3">
        </div>
        <div class="grid grid-cols-3 gap-8 neko-wrap">
        {{range .SudokuGenerates}}
          <a href="/sudoku/{{.SudokuGenerate.Id}}" >
            {{if eq .SudokuGenerate.Level 1}}
            <div class="box neko-box-a bg-gray-100">
              <div class="px-3 mt-2 pb-2">
                <div class="rounded-full bd-blue-100 text-neutral-700 hover:text-neutral-800">
                  ★★★Easy
                </div>
              </div>
              <div class="text-neutral-700 hover:text-neutral-800">
                NO.{{.SudokuGenerate.Id}}
              </div>
              <div class="text-neutral-700 hover:text-neutral-800">
                進捗：{{.Progress.ProgressRate}}%
              </div>
            </div>
            {{end}}
            {{if eq .SudokuGenerate.Level 2}}
            <div class="box neko-box-d bg-gray-100">
              <div class="px-3 mt-2 pb-2">
                <div class="rounded-full bd-blue-200 text-neutral-700 hover:text-neutral-800">
                  ★★☆Nomal
                </div>
              </div>
              <div class="text-neutral-700 hover:text-neutral-800">
                NO.{{.SudokuGenerate.Id}}
              </div>
              <div class="text-neutral-700 hover:text-neutral-800">
                進捗：{{.Progress.ProgressRate}}%
              </div>
            </div>
            {{end}}
            {{if eq .SudokuGenerate.Level 3}}
            <div class="box neko-box-c bg-gray-100">
              <div class="px-3 mt-2 pb-2">
                <div class="rounded-full bd-red-100 text-neutral-700 hover:text-neutral-800">
                  ★☆☆Hard
                </div>
              </div>
              <div class="text-neutral-700 hover:text-neutral-800">
                NO.{{.SudokuGenerate.Id}}
              </div>
              <div class="text-neutral-700 hover:text-neutral-800">
                進捗：{{.Progress.ProgressRate}}%
              </div>
            </div>
            {{end}}
          </a>
        {{end}}
        </div>
      </div>
    </div>
  </div>
</body>
</html>
