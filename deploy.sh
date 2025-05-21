#!/bin/bash

REPOSITORY_URL="https://github.com/cyworld8x/go-postgres-kubernetes-grpc.git"
BLUE_PATH="/opt/api/blue"
GREEN_PATH="/opt/api/green"
CURRENT_PATH="/opt/api"
APP_NAME="user-api"
LOG_FILE="$BLUE_PATH/deployment.log" # Or a common log file
TIMESTAMP=$(date +"%Y-%m-%d %H:%M:%S")

get_current_active() {
    readlink "$CURRENT_PATH" | sed 's/\/opt\/api\///'
}

get_inactive_color() {
    local active=$(get_current_active)
    if [ "$active" == "blue" ]; then
        echo "green"
    else
        echo "blue"
    fi
}

deploy_to_env() {
    local target_color="$1"
    local target_path="/opt/api/$target_color"

    echo "[$TIMESTAMP] Deploying to $target_color environment ($target_path)..." >> "$LOG_FILE"

    if [ ! -d "$target_path" ]; then
        mkdir -p "$target_path"
    fi

    cd "$target_path" || {
        echo "Error: Could not navigate to $target_path." >> "$LOG_FILE"
        exit 1
    }

    # If it's a fresh deployment to this color, clone the repository
    if [ ! -d ".git" ]; then
        echo "[$TIMESTAMP] Cloning repository..." >> "$LOG_FILE"
        git clone "$REPOSITORY_URL" . >> "$LOG_FILE" 2>&1
    else
        echo "[$TIMESTAMP] Pulling latest changes from Git..." >> "$LOG_FILE"
        git pull origin main >> "$LOG_FILE" 2>&1 # Adjust 'main' to your main branch
    fi

    echo "[$TIMESTAMP] Installing dependencies and building application..." >> "$LOG_FILE"

    go build -o "cmd/user/$APP_NAME" cmd/user/main.go >> "$LOG_FILE" 2>&1
    sudo mkdir -p $target_path/cmd/user/config
    sudo chown  ec2-user:ec2-user $target_path

    scp -o StrictHostKeyChecking=no "cmd/user/$APP_NAME" " $target_path/$APP_NAME" 
    scp -o StrictHostKeyChecking=no "cmd/user/config/app.env" $target_path/cmd/user/config
    # Add your build, install dependencies commands here
    # Example for Node.js:
    # npm install >> "$LOG_FILE" 2>&1
    # npm run build >> "$LOG_FILE" 2>&1
    # Example for Python:
    # pip install -r requirements.txt >> "$LOG_FILE" 2>&1

    echo "[$TIMESTAMP] Deployment to $target_color complete." >> "$LOG_FILE"
}

run_tests() {
    local target_color="$1"
    local target_path="/opt/api/$target_color"

    echo "[$TIMESTAMP] Running tests on $target_color environment..." >> "$LOG_FILE"
    cd "$target_path" || return 1
    # Add your test commands here
    # Example for Node.js:
    # npm test >> "$LOG_FILE" 2>&1
    # If tests fail, exit with an error code
    # if [ $? -ne 0 ]; then
    #     echo "[$TIMESTAMP] Tests failed on $target_color. Aborting switch." >> "$LOG_FILE"
    #     return 1
    # fi
    echo "[$TIMESTAMP] Tests passed on $target_color." >> "$LOG_FILE"
    return 0
}

switch_to_env() {
    local target_color="$1"
    echo "[$TIMESTAMP] Switching traffic to $target_color environment..." >> "$LOG_FILE"
    rm "$CURRENT_PATH"
    ln -s "/opt/api/$target_color" "$CURRENT_PATH"
    sudo systemctl reload nginx # Or your service to reload
    echo "[$TIMESTAMP] Traffic switched to $target_color." >> "$LOG_FILE"
}

inactive_color=$(get_inactive_color)
echo "[$TIMESTAMP] Inactive environment is: $inactive_color" >> "$LOG_FILE"

deploy_to_env "$inactive_color"

echo "[$TIMESTAMP] Running tests on the new deployment ($inactive_color)..." >> "$LOG_FILE"
if run_tests "$inactive_color"; then
    echo "[$TIMESTAMP] Tests passed. Switching to $inactive_color..." >> "$LOG_FILE"
    switch_to_env "$inactive_color"
    echo "[$TIMESTAMP] Blue/Green deployment successful." >> "$LOG_FILE"
else
    echo "[$TIMESTAMP] Tests failed. Not switching traffic." >> "$LOG_FILE"
fi

exit 0