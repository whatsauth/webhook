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
?>
