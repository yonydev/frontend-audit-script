# Go Frontend Audit Script Action

Owned by the Core Team

This GitHub Action runs a Go-based script to audit a frontend repository. The script evaluates various aspects of the repository, such as React version, icon libraries, styling libraries, theme providers, web fonts, and assets. It generates a detailed evaluation report and posts it as a comment on the associated pull request.

## Features

- **React Version Check**: Ensures the project uses a supported React version.
- **Icon Libraries Evaluation**: Checks for common icon libraries and evaluates their usage.
- **Styling Libraries Evaluation**: Evaluates the presence of allowed and disallowed styling libraries.
- **Theme Providers Check**: Identifies theme provider components in the codebase.
- **Web Fonts Check**: Analyzes the usage of web fonts in stylesheets.
- **Assets Optimization Check**: Identifies large assets that may need optimization.
- **Score Calculation**: Provides a normalized score (0-10) based on the evaluations.

## Inputs

- `token` (required): GitHub token to add comments to the pull request.

## Outputs

The action posts a detailed evaluation report as a comment on the pull request, including scores for each evaluation and a total normalized score.

## Example Usage

> [!NOTE]  
> If you are auditing a **monorepo**, the action does not fully support it yet.

> [!IMPORTANT]  
> It is recommended to put this action after the build step in your workflow to
> ensure that the code is built and ready for evaluation.

Below is an example of how to use this action in a GitHub Actions workflow:

```yml
name: Frontend Audit

on:
  pull_request:
    branches:
      - main

jobs:
  audit:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v3

      - name: Run Frontend Audit Script
        uses: yonatanrojas/frontendauditscript@master
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
```

## How It Works

The action runs the Go script inside a Docker container.
The script evaluates the repository based on predefined criteria.
The results are posted as a comment on the pull request, providing insights into the repository's quality and areas for improvement.

## 📊 How the Scoring System Works

The **Go Frontend Audit Script Action** evaluates different areas of a frontend project (like React version, libraries used, styling, etc.) and assigns scores based on quality and best practices.

Each evaluation includes:

- **Score**: How well this part of the code is doing.
- **Min/Max Score**: The possible score range for that evaluation.
- **Weight**: How important this evaluation is in the overall score.

---

## 🧮 How the Final Score is Calculated

1. **Weighted Score**: Each evaluation's score is multiplied by its weight.
2. **Range Total**: We calculate the total possible minimum and maximum weighted scores.
3. **Normalize**: The result is scaled to a consistent range of **0 to 10**, so different evaluations can be fairly compared.

### 🔁 Example

- React Version Score: 2 (out of -2 to 2), weighted 3 → contributes **6**
- Icon Libraries Score: -2 (out of -3 to 0), weighted 2 → contributes **-4**
- Styling Libraries Score: 1 (out of -3 to 2), weighted 4 → contributes **4**

**Total weighted score**: `6 - 4 + 4 = 6`  
**Possible range**: `-20 to 14`  
**Normalized Score**: **7.65 out of 10**

---

## ✅ Key Takeaways

- Evaluations with **higher weights affect the final score more**.
- All scores are **normalized to a 0–10 scale**.
- Final results are made available via environment variables for use in CI/CD pipelines or dashboards.

This system ensures scores are **fair**, **balanced**, and **easy to understand**, no matter how different the evaluations are.
