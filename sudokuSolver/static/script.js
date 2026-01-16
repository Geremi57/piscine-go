document.addEventListener("DOMContentLoaded", () => {
  const tableBody = document.querySelector("#sudokuTable tbody");

  for (let i = 0; i < 9; i++) {
    let row = document.createElement("tr");
    for (let j = 0; j < 9; j++) {
      let cell = document.createElement("td");
      let input = document.createElement("input");
      input.setAttribute("maxlength", "1");
      input.dataset.row = i;
      input.dataset.col = j;
      cell.appendChild(input);
      row.appendChild(cell);
    }
    tableBody.appendChild(row);
  }

  const inputs = document.querySelectorAll("#sudokuTable input");
  const form = document.getElementById("sudokuForm");
  const errorDiv = document.getElementById("error");
  const clearBtn = document.getElementById("clearBtn");

  inputs.forEach((inp) => {
    inp.addEventListener("keydown", (e) => {
      let row = parseInt(inp.dataset.row);
      let col = parseInt(inp.dataset.col);

      switch (e.key) {
        case "ArrowUp":
          if (row > 0) focusCell(row - 1, col);
          e.preventDefault();
          break;
        case "ArrowDown":
          if (row < 8) focusCell(row + 1, col);
          e.preventDefault();
          break;
        case "ArrowLeft":
          if (col > 0) focusCell(row, col - 1);
          e.preventDefault();
          break;
        case "ArrowRight":
          if (col < 8) focusCell(row, col + 1);
          e.preventDefault();
          break;
      }
    });
  });

  function focusCell(r, c) {
    const nextInput = document.querySelector(
      `#sudokuTable input[data-row="${r}"][data-col="${c}"]`
    );
    if (nextInput) nextInput.focus();
  }

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    let board = [];
    let inputs = document.querySelectorAll("#sudokuTable input");

    for (let i = 0; i < 9; i++) {
      let row = [];
      for (let j = 0; j < 9; j++) {
        let val = inputs[i * 9 + j].value.trim();
        row.push(val);
      }
      board.push(row);
    }

    try {
      let res = await fetch("http://localhost:8080/solve", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(board),
      });

      if (!res.ok) {
        let errText = await res.text();
        errorDiv.innerText = errText;
        return;
      }

      let solved = await res.json();
      errorDiv.innerText = "";

      solved.flat().forEach((val, idx) => {
        inputs[idx].value = val;
      });
    } catch (err) {
      errorDiv.innerText = "Server error. Please try again.";
    }
  });

  clearBtn.addEventListener("click", () => {
    document
      .querySelectorAll("#sudokuTable input")
      .forEach((inp) => (inp.value = ""));
    errorDiv.innerText = "";
  });
});
