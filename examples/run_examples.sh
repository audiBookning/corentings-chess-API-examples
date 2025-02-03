#!/bin/bash

# Create a logs directory if it doesn't exist
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
logs_dir="$script_dir/logs"
mkdir -p "$logs_dir"

# Get current timestamp for the log files
timestamp=$(date +"%Y-%m-%d_%H-%M-%S")

# Find all main.go files in the examples subdirectories
main_files=$(find "$script_dir" -name "main.go")
file_count=$(echo "$main_files" | wc -l)

echo "Found $file_count example programs to run"
echo

# Function to print colored output
print_success() {
    echo -e "\033[0;32m✓\033[0m $1"
}

print_error() {
    echo -e "\033[0;31m✗\033[0m $1"
}

while IFS= read -r main_file; do
    # Get example name from parent directory
    example_name=$(basename "$(dirname "$main_file")")
    echo "Running example: $example_name"
    
    # Create log file name
    log_file="$logs_dir/${example_name}_${timestamp}.log"
    
    # Change to the directory containing main.go
    pushd "$(dirname "$main_file")" > /dev/null
    
    # Run the example and capture output
    if output=$(go run main.go 2>&1); then
        # Save output to log file
        echo "$output" > "$log_file"
        print_success "Completed. Log saved to: $log_file"
    else
        print_error "Error running example"
        echo "Error running example: $output" > "$log_file"
    fi
    echo
    
    # Return to original directory
    popd > /dev/null
done <<< "$main_files"

echo "All examples completed. Logs are available in the $logs_dir directory."
