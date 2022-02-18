function usage_pricing_header_table()
  headers = {
    'Nylas-Namespace-Public-Id',
    'Nylas-Namespace-Id',
    'Nylas-Http-Request',
    'Nylas-Application-Public-Id',
    'Nylas-Application-Id',
    'Nylas-Endpoint',
    'Nylas-Request-Uid',
    'Nylas-Event',
    'Nylas-Provider-Name',
    'Nylas-Result-Count',
  }
  result = {}
  for _, header_key in ipairs(headers) do
    ctx_key = string.gsub(header_key:lower(), '-', '_')
    result[header_key] = ctx_key
  end
  return result
end

result = usage_pricing_header_table()
for header_key, ctx_key in pairs(usage_pricing_header_table()) do
  print(header_key, ctx_key)
end
