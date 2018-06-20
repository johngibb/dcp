NUM=$1
if [[ $NUM == "" ]]; then
    echo "usage: ./new.sh number"
    exit 1
fi

set -e

if [[ "$NUM" =~ "^[0-9+]$" ]]; then
    NUM=$(printf '%03d' $NUM)
    NAME="prob${NUM}"
else
    NAME="$NUM"
fi

mkdir "$NAME"

cat > "$NAME/prob.go" <<EOF
package prob
EOF

cat > "$NAME/prob_test.go" <<EOF
package prob

import (
    "testing"
)

func Test(t *testing.T) {
}
EOF
