NUM=$1
if [[ $NUM == "" ]]; then
    echo "usage: ./new.sh number"
    exit 1
fi

set -e

NUM=$(printf '%03d' $NUM)
NAME="prob${NUM}"

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
