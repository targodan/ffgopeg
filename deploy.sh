#!/bin/bash

repo="$(git config remote.origin.url)"
repo="${repo/https:\/\/github.com\//git@github.com:}"
commit=$(git rev-parse --verify HEAD)
branch="$TRAVIS_BRANCH"
user="$DEPLOY_GIT_NAME"
email="$DEPLOY_GIT_EMAIL"
commitMessage="$(echo $DEPLOY_COMMIT_MSG | sed "s/COMMIT_SHA/$commit/g" | sed "s/BRANCH/$branch/g")"

echo "---- DEBUG ----"
echo "repo = $repo"
echo "commit = $commit"
echo "user = $user"
echo "email = $email"
echo "commitMessage = $commitMessage"
echo "branch = $branch"
echo "---------------"

echo "Starting deploy script on branch \"$branch\"."

if [[ "$branch" != "master" || "$ffmpeg" != "3.1.5" || "`go version | grep go1.7`" == "" ]]; then
    echo "Wrong branch, wont deploy."
    exit
fi

echo "Decrypting deploy key..."
openssl aes-256-cbc -K $encrypted_db86d75584f0_key -iv $encrypted_db86d75584f0_iv -in deployKey.enc -out deployKey -d && echo "Done." || echo "Error decoding!"
chmod 600 deployKey
eval "$(ssh-agent -s)"
ssh-add deployKey

repoDir=`pwd`
cd $HOME

echo "Cloning gh-pages..."
git clone --branch=gh-pages "$repo" pagesOut && echo "Done." || echo "Error cloning."
cd pagesOut

git config user.name "$user"
git config user.email "$email"

echo "Generating searcher info..."
go run generateData.go "$repoDir" > data.js && echo "Done." || echo "Error generating."

echo "Committing..."
git add -A
git commit -m "$commitMessage" && echo "Done." || echo "Error committing."

echo "Pushing..."
git push origin gh-pages && echo "Done." || echo "Error pushing."

echo "Deploy script finished."
