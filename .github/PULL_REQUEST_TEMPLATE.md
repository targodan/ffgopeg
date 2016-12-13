# Before submitting

Please make sure all boxes are ticked before submitting.

- [ ] The base branch of your changes is the `develop` branch. Use `git rebase --onto develop <old-parent>` to fix it.
- [ ] Your code is go fmt'd.
- [ ] Github says it can be merged without any conflicts.
- [ ] You successfully compiled on your machine.
- [ ] All tests were successfull on your machine.
- [ ] You read this *whole* text and deleted everything except this checklist.

# After submitting

- Please wait for TravisCI build and test your request (this may take some time, feel free to come back tomorrow or even next week). If any issues occur please fix them by adding new commits to your branch. (Travis should then automatically rebuild.)
- If your branch history is messy please do an interactive rebase in order to clean up. (`git rebase -i HEAD~n`) Your pull request may include multiple commits but they should each contain *significant* changes or fixes.

# After TravisCI was successfull

A project maintainer will review your code and may ask you to make some last minute changes.

NOTE: The project maintainer(s) will ignore any requests until they were successfully built by TravisCI.

# After successfull code review

Your request will be merged.
Thank you for contributing! :smiley:
