* 如何修改git提交历史里的邮箱

git config alias.change-commits '!'"f() { VAR=\$1; OLD=\$2; NEW=\$3; shift 3; git filter-branch --env-filter \"if [[ \\\"\$\`echo \$VAR\`\\\" = '\$OLD' ]]; then export \$VAR='\$NEW'; fi\" \$@; }; f "
git change-commits GIT_AUTHOR_NAME erxiangbo@wacai.com shijian0912@163.com -f
git change-commits GIT_AUTHOR_EMAIL <old@email.com> <new@email.com> -f
git change-commits GIT_COMMITTER_NAME "<Old Name>" "<New Name>" -f
git change-commits GIT_COMMITTER_EMAIL <old@email.com> <new@email.com> -f