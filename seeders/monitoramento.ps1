# PowerShell script to mimic monitoramento.sh behavior on Windows
$ErrorActionPreference = 'Stop'
Set-StrictMode -Version Latest

$ApiBase = 'http://localhost:8181/api/v1'
$ProfessionalEmail = 'joao.silva@mindtrace.com'
$ProfessionalPassword = 'Password123!'
$Pacientes = @(
    'ana.paciente@mindtrace.com',
    'bruno.paciente@mindtrace.com',
    'carla.paciente@mindtrace.com',
    'diego.paciente@mindtrace.com'
)
$PacienteIds = @(1, 2, 3, 4)

function Invoke-ApiLogin {
    param (
        [string]$Email,
        [string]$Senha
    )

    $payload = @{ email = $Email; senha = $Senha } | ConvertTo-Json -Compress

    try {
        $response = Invoke-RestMethod -Method Post -Uri "$ApiBase/entrar/login" -ContentType 'application/json' -Body $payload -ErrorAction Stop
    } catch {
        Write-Warning "Login request failed for ${Email}: $_"
        return $null
    }

    if (($response | Get-Member -MemberType All -Name "dados") -and ($response | Get-Member -MemberType All -Name "dados.token")) { return $response.dados.token }
    if  (($response | Get-Member -MemberType All -Name "token")) { return $response.token }

    return $null
}

function Invoke-PatientMonitoring {
    param (
        [string]$Token,
        [int]$PacienteId,
        [int]$Registros
    )

    $headers = @{ Authorization = "Bearer $Token" }
    $uri = "$ApiBase/monitoramento/paciente-lista?pacienteID=$PacienteId&numRegistros=$Registros"

    try {
        return Invoke-RestMethod -Method Get -Uri $uri -Headers $headers -ContentType 'application/json' -ErrorAction Stop
    } catch {
        Write-Warning "Monitoring request failed for paciente ${PacienteId}: $_"
        return $null
    }
}

Write-Host 'Starting monitoring tests...' -ForegroundColor Cyan

$professionalToken = Invoke-ApiLogin -Email $ProfessionalEmail -Senha $ProfessionalPassword
if (-not $professionalToken) {
    Write-Error 'Professional login failed. Aborting.'
    exit 1
}

Write-Host "Professional token acquired (first 30 chars): $($professionalToken.Substring(0, [Math]::Min(30, $professionalToken.Length)))"

for ($i = 0; $i -lt 4; $i++) {
    $pacEmail = $Pacientes[$i]
    $pacId = $PacienteIds[$i]

    Write-Host "Monitoring $pacEmail (ID: $pacId)..." -ForegroundColor Yellow
    $monitoramento = Invoke-PatientMonitoring -Token $professionalToken -PacienteId $pacId -Registros 4

    if (-not $monitoramento) {
        Write-Warning 'Monitoring request failed.'
        continue
    }

    Write-Host 'Monitoring response:'
    $monitoramento | ConvertTo-Json -Depth 6

    $dados = $monitoramento.dados_monitoramento
    if ($dados -and $dados.Count -gt 0) {
        Write-Host 'Records:'
        foreach ($registro in $dados) {
            Write-Host ("  Date: {0} | Humor: {1} | Sleep: {2}h | Energy: {3} | Stress: {4}" -f `
                $registro.data, $registro.nivel_humor, $registro.horas_sono, $registro.nivel_energia, $registro.nivel_stress)
        }
    } else {
        Write-Host 'No monitoring records for this patient.'
    }

    Write-Host '---'
}

Write-Host 'Monitoring tests completed.' -ForegroundColor Green
