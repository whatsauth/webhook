<?php
// Define your secret token here
define('SECRET_TOKEN', 'your_secret_token');

// Check if the request method is POST
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    // Get the token from the request headers
    $headers = apache_request_headers();
    $requestToken = isset($headers['Request-Token']) ? $headers['Request-Token'] : '';

    // Verify the token
    if ($requestToken === SECRET_TOKEN) {
        // Get the JSON payload from the POST request
        $jsonPayload = file_get_contents('php://input');

        // Decode the JSON payload
        $data = json_decode($jsonPayload, true);

        // Check if JSON decoding was successful
        if ($data) {
            // Use $data as an associative array with your WAMessage struct fields
            // For example, access phone_number like $data['phone_number']
            // Here, you would process the data or save it to a database, etc.

            // Send a response back
            echo json_encode(['status' => 'success', 'message' => 'Data received and processed']);
        } else {
            // JSON was not decoded, handle the error
            http_response_code(400);
            echo json_encode(['status' => 'error', 'message' => 'Invalid JSON']);
        }
    } else {
        // Token is invalid, handle the error
        http_response_code(401);
        echo json_encode(['status' => 'error', 'message' => 'Unauthorized access']);
    }
} else {
    // The request method is not POST, handle the error
    http_response_code(405);
    echo json_encode(['status' => 'error', 'message' => 'Method not allowed']);
}


// The JSON data.
$data = array(
    'to' => '6281111',
    'isgroup' => false,
    'messages' => 'hai kak'
);

// JSON encode the data.
$jsonData = json_encode($data);

// The URL to send the POST request to.
$url = 'http://example.com/endpoint'; // Replace with the URL of the server you're sending the request to.

// The token to be sent in the header.
$token = 'YourTokenHere'; // Replace this with your actual token.

// Initialize cURL session.
$ch = curl_init($url);

// Set cURL options.
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $jsonData);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, array(
    'Content-Type: application/json',
    'Request-Token: ' . $token // Add your token header here
));

// Execute the POST request.
$response = curl_exec($ch);

// Check for cURL errors
if (curl_errno($ch)) {
    $error_msg = curl_error($ch);
}

// Close cURL session.
curl_close($ch);

// If there was an error, handle it here.
if (isset($error_msg)) {
    // Log or echo the error message.
    echo "cURL Error: " . $error_msg;
}

// Do something with the response.
echo $response;
?>