// package: heyrenee.v1
// file: heyrenee/v1/diagnosis_service.proto

var heyrenee_v1_diagnosis_service_pb = require("../../heyrenee/v1/diagnosis_service_pb");
var heyrenee_v1_messages_patient_diagnosis_pb = require("../../heyrenee/v1/messages/patient_diagnosis_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var DiagnosisService = (function () {
  function DiagnosisService() {}
  DiagnosisService.serviceName = "heyrenee.v1.DiagnosisService";
  return DiagnosisService;
}());

DiagnosisService.DiagnosisSuggest = {
  methodName: "DiagnosisSuggest",
  service: DiagnosisService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_diagnosis_service_pb.DiagnosisSuggestRequest,
  responseType: heyrenee_v1_diagnosis_service_pb.DiagnosisSuggestResponse
};

DiagnosisService.CreatePatientDiagnosis = {
  methodName: "CreatePatientDiagnosis",
  service: DiagnosisService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_diagnosis_service_pb.CreatePatientDiagnosisRequest,
  responseType: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis
};

DiagnosisService.UpdatePatientDiagnosis = {
  methodName: "UpdatePatientDiagnosis",
  service: DiagnosisService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_diagnosis_service_pb.UpdatePatientDiagnosisRequest,
  responseType: heyrenee_v1_messages_patient_diagnosis_pb.PatientDiagnosis
};

DiagnosisService.ListPatientDiagnoses = {
  methodName: "ListPatientDiagnoses",
  service: DiagnosisService,
  requestStream: false,
  responseStream: false,
  requestType: heyrenee_v1_diagnosis_service_pb.ListPatientDiagnosesRequest,
  responseType: heyrenee_v1_diagnosis_service_pb.ListPatientDiagnosesResponse
};

exports.DiagnosisService = DiagnosisService;

function DiagnosisServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

DiagnosisServiceClient.prototype.diagnosisSuggest = function diagnosisSuggest(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(DiagnosisService.DiagnosisSuggest, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DiagnosisServiceClient.prototype.createPatientDiagnosis = function createPatientDiagnosis(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(DiagnosisService.CreatePatientDiagnosis, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DiagnosisServiceClient.prototype.updatePatientDiagnosis = function updatePatientDiagnosis(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(DiagnosisService.UpdatePatientDiagnosis, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

DiagnosisServiceClient.prototype.listPatientDiagnoses = function listPatientDiagnoses(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(DiagnosisService.ListPatientDiagnoses, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.DiagnosisServiceClient = DiagnosisServiceClient;

