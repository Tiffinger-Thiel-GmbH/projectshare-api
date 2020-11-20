import React, { useEffect, useState } from 'react';
import ProjectTable from '../components/ProjectTable';
import DocumentTable from '../components/DocumentTable';
import { Grid, Container, Box, TextField, makeStyles, Button, IconButton } from '@material-ui/core';
import { DocumentDTO, getDocumentsByProjectId, getProjects, ProjectDTO, uploadDocument, createProject } from '../api/apis';
import AddCircleOutlineIcon from '@material-ui/icons/AddCircleOutline';
import CloudUploadIcon from '@material-ui/icons/CloudUpload';
const useStyles = makeStyles(theme => ({
  margin: {
    margin: theme.spacing(1)
  }
}));

export default function Content() {
  const [selectedProject, setSelectedProject] = useState<string>();
  const [projects, setProjects] = useState<ProjectDTO[]>([]);
  const [documents, setDocuments] = useState<DocumentDTO[]>([]);
  const [documentToUpload, setDocumentToUpload] = useState<File | undefined>();
  const [projectToCreate, setProjectToCreate] = useState('');
  const style = useStyles();
  const fileUpload = React.createRef<HTMLInputElement>();

  useEffect(() => {
    getProjects().then(setProjects);
  }, [projectToCreate]);

  useEffect(() => {
    if (selectedProject) getDocumentsByProjectId(selectedProject).then(setDocuments);
  }, [selectedProject, documentToUpload]);

  const onUploadDocument = async (file?: File | null) => {
    if (file && selectedProject) {
      const newFormData = new FormData();

      newFormData.append('file', file, file.name);

      await uploadDocument(selectedProject, newFormData).then(() => setDocumentToUpload(file));
    }
  };

  const onCreateProject = async () => {
    if (projectToCreate !== '') {
      await createProject(projectToCreate).then(() => setProjectToCreate(''));
    }
  };

  return (
    <>
      <Container maxWidth="xl">
        <Grid container justify="center" spacing={4}>
          <Grid item xs={6}>
            <Box ml={8} mt={4} textAlign="center">
              <div className={style.margin}>
                <Grid container spacing={1} alignItems="flex-end" style={{ minHeight: '56px' }}>
                  <Grid item xs={1}>
                    <IconButton onClick={() => onCreateProject()}>
                      <AddCircleOutlineIcon />
                    </IconButton>
                  </Grid>
                  <Grid item xl={11}>
                    <TextField
                      id="input-with-icon-grid"
                      label="New Project"
                      fullWidth
                      value={projectToCreate}
                      onChange={e => setProjectToCreate(e.target.value)}
                    />
                  </Grid>
                </Grid>
              </div>
              <ProjectTable setSelectedProject={setSelectedProject} projects={projects} />
            </Box>
          </Grid>
          <Grid item xs={6}>
            <Box mr={8} mt={4} textAlign="center">
              <div className={style.margin}>
                <Grid container spacing={1} alignItems="flex-end" style={{ minHeight: '56px' }}>
                  <Grid item xs={12}>
                    <Button onClick={() => fileUpload.current?.click()} variant="contained" color="default" startIcon={<CloudUploadIcon />}>
                      Upload
                    </Button>
                    <input
                      type="file"
                      style={{ display: 'none ' }}
                      onChange={e => onUploadDocument(e.target.files?.item(0))}
                      ref={fileUpload}
                    />
                  </Grid>
                </Grid>
              </div>
              <DocumentTable documents={documents} />
            </Box>
          </Grid>
        </Grid>
      </Container>
    </>
  );
}
