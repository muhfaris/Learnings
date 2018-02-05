## Replace master with other branch
`git checkout seotweaks
git merge -s ours master
git checkout master
git merge seotweaks
 `
(`-s` ours is short for `--strategy=ours`)
 

