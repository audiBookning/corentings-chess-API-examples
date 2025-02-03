# Create a logs directory if it doesn't exist
$scriptDir = Split-Path -Parent $MyInvocation.MyCommand.Path
$logsDir = Join-Path $scriptDir "logs"
if (-not (Test-Path $logsDir)) {
    New-Item -ItemType Directory -Path $logsDir
}

# Get current timestamp for the log files
$timestamp = Get-Date -Format "yyyy-MM-dd_HH-mm-ss"

# Find all main.go files in the examples subdirectories
$mainFiles = Get-ChildItem -Path $scriptDir -Recurse -Filter "main.go"

Write-Host "Found $($mainFiles.Count) example programs to run`n"

foreach ($mainFile in $mainFiles) {
    $exampleName = Split-Path (Split-Path $mainFile.FullName -Parent) -Leaf
    Write-Host "Running example: $exampleName"
    
    # Create log file name
    $logFile = Join-Path $logsDir "$($exampleName)_$timestamp.log"
    
    # Change to the directory containing main.go
    Push-Location (Split-Path $mainFile.FullName -Parent)
    
    try {
        # Run the example and capture output
        $output = & go run main.go 2>&1
        
        # Save output to log file
        $output | Out-File -FilePath $logFile -Encoding UTF8
        
        Write-Host "✓ Completed. Log saved to: $logFile`n"
    }
    catch {
        Write-Host "✗ Error running example: $_`n" -ForegroundColor Red
        "Error running example: $_" | Out-File -FilePath $logFile -Encoding UTF8
    }
    finally {
        # Return to original directory
        Pop-Location
    }
}

Write-Host "All examples completed. Logs are available in the $logsDir directory."
