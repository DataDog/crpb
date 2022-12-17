var builder = WebApplication.CreateBuilder(args);

var port = Environment.GetEnvironmentVariable("PORT") ?? "8080";
var url = $"http://0.0.0.0:{port}";

var app = builder.Build();

var target = Environment.GetEnvironmentVariable("TARGET") ?? "World";

app.MapGet("/", () => $"Hello {target} from .NET 6.0!");

app.Run(url);