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

**Important**:

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
        uses: yonatanrojas/frontendauditscript@latest
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
```

## How It Works

The action runs the Go script inside a Docker container.
The script evaluates the repository based on predefined criteria.
The results are posted as a comment on the pull request, providing insights into the repository's quality and areas for improvement.
