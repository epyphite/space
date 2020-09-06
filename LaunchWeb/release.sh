TIME=$2
VERSION=$1
TOOLS=tools
LIB=LIB
BIN=bin
DEST=build/dist
TAG=$VERSION
DESCRIPTION="Automated Distribution"
RELEASENAME=$VERSION


echo "Uploading to GITHUB"


github-release release \
    --user epyphite \
    --repo website \
    --tag $TAG \
    --name "$RELEASENAME" \
    --description "$DESCRIPTION" \
    --pre-release

github-release upload \
    --user epyphite \
    --repo website \
    --tag $TAG  \
    --name linux-dist.tar.gz \
    --file $DEST/linux-dist.tar.gz

github-release upload \
    --user epyphite \
    --repo website \
    --tag $TAG  \
    --name linux-dist.zip \
    --file $DEST/linux-dist.zip
