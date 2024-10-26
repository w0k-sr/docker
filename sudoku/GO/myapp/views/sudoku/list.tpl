<!DOCTYPE html>

<html>
<head>
  <title>数独一覧</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Hachi+Maru+Pop&display=swap" rel="stylesheet">
  <link rel="stylesheet" href="../static/css/modern-normalize.css">
  <link rel="stylesheet" type="text/css" href="../static/css/main.css">
  <script src="../static/js/sudoku/list.js"></script>
  <script src="https://code.jquery.com/jquery-3.7.1.slim.min.js" 
    integrity="sha256-kmHvs0B+OpCW5GVHUNjv9rOmY0IvSIRcf7zGUDTDQM8=" crossorigin="anonymous"></script>
</head>

<body>
  <div class="min-h-screen bg-neutral-300 relative">
    <div class="pt-16 md:pt-20 pb-28 container px-2 md:px-5 xl:px-20">
      <div class="m-auto">
        <div class="px-2 lg:px-10 rounded-t-md bg-neutral-600 flex flex-wrap flex-space">
          <h1 class="text-neutral-100">数独一覧</h1>
            <div class="text-neutral-100 items-center" style="font-size: 2rem;">
              {{.Message}}
            </div>
            <a href="/sudoku/create" class="btn btn-custom01">
              <span class="btn-custom01-front">問題を作成</span>
            </a>
        </div>
        <div class="bg-gray-300 min-h-screen">

          <div class="grid grid-cols-3 gap-1 neko-wrap">
          {{range .SudokuGenerates}}
            <a href="/sudoku/{{.SudokuGenerate.Id}}" >
              {{if eq .SudokuGenerate.Level 1}}
              <div class="box neko-box-a bg-gray-100">
                <div class="px-3 pb-2 flex flex-wrap  flex-space">
                  <div class="py-2 px-3 rounded-full bd-blue-100 text-neutral-700 hover:text-neutral-800">
                    <span style="color: #ffc107;">★★★</span><br> Easy
                  </div>
                  <div class="text-neutral-700 hover:text-neutral-800">
                    NO.{{.SudokuGenerate.Id}}
                  </div>
                </div>
                <div class="text-neutral-700 hover:text-neutral-800">
                  進捗：{{.Progress.ProgressRate}}%
                </div>
              </div>
              {{end}}
              {{if eq .SudokuGenerate.Level 2}}
              <div class="box neko-box-d bg-gray-100">
                <div class="px-3 pb-2 flex flex-wrap  flex-space">
                  <div class="py-2 px-3 rounded-full bd-blue-200 text-neutral-700 hover:text-neutral-800">
                    <span style="color: #ffc107;">★★☆</span><br> Nomal
                  </div>
                  <div class="text-neutral-700 hover:text-neutral-800">
                    NO.{{.SudokuGenerate.Id}}
                  </div>
                </div>
                <div class="text-neutral-700 hover:text-neutral-800">
                  進捗：{{.Progress.ProgressRate}}%
                </div>
              </div>
              {{end}}
              {{if eq .SudokuGenerate.Level 3}}
              <div class="box neko-box-c bg-gray-100">
                <div class="px-3 pb-2 flex flex-wrap  flex-space">
                  <div class="py-2 px-3 rounded-full bd-red-100 text-neutral-700 hover:text-neutral-800">
                    <span style="color: #ffc107;">★☆☆</span><br> Hard
                  </div>
                  <div class="text-neutral-700 hover:text-neutral-800">
                    NO.{{.SudokuGenerate.Id}}
                  </div>
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
  </div>
</body>
</html>
