// initializing elastic search client
const esClientObj = ElasticSearchClientGh.getInstance();

/**
 * Fetches library records from DynamoDB based on the provided library names.
 * @param libNames - An array of library names.
 * @returns A promise that resolves to an array of LibraryRecord objects.
 * @throws If there is an error fetching the items from DynamoDB.
 */
