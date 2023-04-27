const https = require('https');

// Function to extract IP address in words from a given string
function extractIpAddress(str) {
    const ipRegex = /(\d{1,3}\s*){4}/g;
    const matches = str.match(ipRegex);
    if (matches) {
        // Remove extra spaces and replace "point" with "."
        const ipStr = matches[0].replace(/\s+/g, '').replace(/point/g, '.');
        return ipStr;
    }
    return null;
}

// Make API call to get the passage with hidden IP address
https.get('https://quest.squadcast.tech/api/RA2011003020090/worded_ip', (res) => {
    let data = '';
    res.on('data', (chunk) => {
        data += chunk;
    });
    res.on('end', () => {
        const ipAddress = extractIpAddress(data);
        console.log("hi");
        if (ipAddress) {
            console.log(ipAddress); // Print the extracted IP address
            // Make submission to the API
            const submissionUrl = `https://quest.squadcast.tech/api/RA2011003020090/submit/worded_ip?answer=${ipAddress}`;
            https.post(submissionUrl, (res) => {
                console.log(`Submission status: ${res.statusCode}`);
            });
        } else {
            console.log('No IP address found in the passage.');
        }
    });
}).on('error', (err) => {
    console.log(`Error getting passage: ${err}`);
});