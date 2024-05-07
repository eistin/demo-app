<!DOCTYPE html>
<html>
<head>
    <title>Counter App</title>
    <link rel="stylesheet" type="text/css" href="styles.css">
    <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js"></script>
    <script>
        $(document).ready(function() {
            // Define API URL from environment variable
            var apiUrl = "<?php echo getenv('API_URL'); ?>";

            // Function to handle incrementing the counter
            $("#incrementBtn").click(function() {
                $.ajax({
                    url: apiUrl + "/increment",
                    type: "POST",
                    success: function(data) {
                        // Refresh the current number after successful increment
                        getCurrentNumber();
                    },
                    error: function(xhr, status, error) {
                        console.error("Error incrementing counter:", error);
                    }
                });
            });

            // Function to handle decrementing the counter
            $("#decrementBtn").click(function() {
                $.ajax({
                    url: apiUrl + "/decrement",
                    type: "POST",
                    success: function(data) {
                        // Refresh the current number after successful decrement
                        getCurrentNumber();
                    },
                    error: function(xhr, status, error) {
                        console.error("Error decrementing counter:", error);
                    }
                });
            });

            // Function to fetch and display the current number
            function getCurrentNumber() {
                $.ajax({
                    url: apiUrl + "/current",
                    type: "GET",
                    success: function(data) {
                        $("#currentNumber").text("Current Number: " + data);
                    },
                    error: function(xhr, status, error) {
                        console.error("Error fetching current number:", error);
                    }
                });
            }

            // Initial fetch of current number when the page loads
            getCurrentNumber();
        });
    </script>
</head>
<body>
    <h1>Counter App</h1>
    <p id="currentNumber">Current Number: </p>
    <button id="incrementBtn">Increment</button>
    <button id="decrementBtn">Decrement</button>
</body>
</html>