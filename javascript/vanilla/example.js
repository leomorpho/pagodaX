window.initializeJS = function initializeApp(targetElement) {
  const container = targetElement || document;

  // Initialize the quiz
  initializeQuiz(container);
};

// Initializes the Quiz
function initializeQuiz(container) {
  // Quiz data
  const questions = [
    {
      question: "What is the capital of France?",
      options: ["Paris", "London", "Berlin", "Madrid"],
      answer: "Paris",
    },
    {
      question: "What is 2 + 2?",
      options: ["3", "4", "5", "6"],
      answer: "4",
    },
  ];

  // Find quiz container
  const quizContainer = container.querySelector("#js-quiz-container");
  if (!quizContainer) return;

  quizContainer.className =
    "max-w-xl mx-auto bg-white shadow-md rounded-lg p-8 mt-5";

  // Render questions
  questions.forEach((q, index) => {
    const questionEl = document.createElement("div");
    questionEl.className = "mb-4 last:mb-0";

    const questionText = document.createElement("p");
    questionText.textContent = index + 1 + ". " + q.question;
    questionText.className = "font-semibold text-lg mb-2";
    questionEl.appendChild(questionText);

    const optionsContainer = document.createElement("div");
    optionsContainer.className = "pl-4";

    q.options.forEach((option) => {
      const optionContainer = document.createElement("div");
      optionContainer.className = "flex items-center mb-1";

      const radioButton = document.createElement("input");
      radioButton.setAttribute("type", "radio");
      radioButton.setAttribute("name", "question" + index);
      radioButton.setAttribute("value", option);
      radioButton.className = "option-input mr-2";

      const label = document.createElement("label");
      label.textContent = option;
      label.className = "select-none";

      optionContainer.appendChild(radioButton);
      optionContainer.appendChild(label);
      optionsContainer.appendChild(optionContainer);
    });

    questionEl.appendChild(optionsContainer);
    quizContainer.appendChild(questionEl);
  });

  // Submit button
  const submitBtn = document.createElement("button");
  submitBtn.textContent = "Submit Answers";
  submitBtn.className =
    "mt-4 px-4 py-2 bg-blue-500 hover:bg-blue-700 text-white font-bold rounded cursor-pointer";
  submitBtn.addEventListener("click", () => {
    let score = 0;
    questions.forEach((q, index) => {
      const selectedOption = container.querySelector(
        `input[name="question${index}"]:checked`
      );
      if (selectedOption && selectedOption.value === q.answer) {
        score++;
      }
    });
    alert(`You scored ${score}/${questions.length}`);
  });

  quizContainer.appendChild(submitBtn);
}
