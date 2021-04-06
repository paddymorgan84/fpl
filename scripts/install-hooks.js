var fs = require("fs");

if (fs.existsSync(".git/hooks") && !process.env.CI) {
  fs.copyFileSync("scripts/hooks/pre-commit", ".git/hooks/pre-commit");
  fs.copyFileSync("scripts/hooks/post-checkout", ".git/hooks/post-checkout");
  fs.copyFileSync("scripts/hooks/post-merge", ".git/hooks/post-merge");
}
