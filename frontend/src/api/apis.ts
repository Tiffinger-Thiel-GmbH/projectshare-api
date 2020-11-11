import { get } from './common';

export interface ProjectDTO {
  id: string;
  name: string;
}

export interface DocumentDTO {
  id: string;
  bucket: string;
  name: string;
}

const baseUrl = 'http://localhost:7777';

export const getProjects = async (): Promise<ProjectDTO[]> => {
  const projects = await get(baseUrl + '/project').then(response => response.data);
  // .catch(() => [
  //   { id: '0975ea14-56a0-4d40-940c-c9e94aa6b359', name: 'ProjectX' },
  //   { id: '22ab3304-5fd8-45ca-b7bd-0ae9836f4e28', name: 'ProjectY' }
  // ]);
  return projects;
};

export const getDocumentsByProjectId = async (projectId: string): Promise<DocumentDTO[]> => {
  const documents = await get(baseUrl + '/project/' + projectId + '/document').then(response => response.data);
  // .catch(() => [{ id: '8eff87f7-6528-48fc-b4ac-e5e22fd195d9', bucket: '22ab3304-5fd8-45ca-b7bd-0ae9836f4e28', name: 'Plan1.txt' }]);

  return documents;
};
